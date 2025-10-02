package functions

import (
	"errors"
	"fmt"
	"github.com/blutspende/go-astm/v3/constants"
	"github.com/blutspende/go-astm/v3/errmsg"
	"github.com/blutspende/go-astm/v3/models/astmmodels"
	"reflect"
)

func ParseStruct(inputLines []string, targetStruct interface{}, lineIndex *int, sequenceNumber int, depth int, config *astmmodels.Configuration) (err error) {
	// Check for maximum depth
	if depth >= constants.MaxDepth {
		return errmsg.ErrStructureParsingMaxDepthReached
	}
	// Check for enough input lines
	if *lineIndex >= len(inputLines) {
		return errmsg.ErrStructureParsingInputLinesDepleted
	}

	// Process the target structure
	targetTypes, targetValues, _, err := ProcessStructReflection(targetStruct)
	if err != nil {
		return err
	}

	// Iterate over the inputFields of the targetStruct struct
	for i, targetType := range targetTypes {
		// Parse the targetStruct field targetFieldAnnotation
		targetStructAnnotation, err := ParseAstmStructAnnotation(targetType)
		if err != nil {
			return err
		}
		// Save the target value pointer
		targetValue := targetValues[i].Addr().Interface()

		// Target is an array it is iterated with conditional break (unknown length)
		if targetStructAnnotation.IsArray {
			// Create the array structure
			sliceType := reflect.SliceOf(targetValues[i].Type().Elem())
			targetValues[i].Set(reflect.MakeSlice(sliceType, 0, 0))

			// Iterate as long as we have matching input structure and still have input lines
			for seq := 1; *lineIndex < len(inputLines); seq++ {
				// Create a new element for the slice to parse into
				elem := reflect.New(targetValues[i].Type().Elem()).Elem()

				nameOk := true
				if targetStructAnnotation.IsComposite {
					// Composite target: recursively parse the composite structure
					err = ParseStruct(inputLines, elem.Addr().Interface(), lineIndex, seq, depth+1, config)
					// If the error is a line type name mismatch, it means the end of the array
					// Note: here an error is used to communicate the end of the array, it is not a real error
					if errors.Is(err, errmsg.ErrStructureParsingLineTypeNameMismatch) {
						nameOk = false
					}
				} else {
					// Non-composite target: parse the line into the new element
					nameOk, err = ParseLine(inputLines[*lineIndex], elem.Addr().Interface(), targetStructAnnotation, seq, config)
					// Increment the line index
					*lineIndex++
				}
				// If the type name is a mismatch, it means the end of the array
				if !nameOk {
					err = nil
					*lineIndex--
					break
				}
				if err != nil {
					return err
				}
				// If no error, add the new element to the slice
				targetValues[i].Set(reflect.Append(targetValues[i], elem))
			}
		} else {
			// Single element structure
			if targetStructAnnotation.IsComposite {
				// Composite target: go further down the rabbit hole
				err = ParseStruct(inputLines, targetValue, lineIndex, 1, depth+1, config)
				if err != nil {
					return err
				}
			} else {
				// Non-composite target: there is a single line to parse
				// Make sure there are enough input lines
				if *lineIndex >= len(inputLines) {
					// Skip if the structure is optional, error otherwise
					if _, exists := targetStructAnnotation.Attributes[constants.AttributeOptional]; exists {
						continue
					} else {
						return errmsg.ErrStructureParsingInputLinesDepleted
					}
				}
				// Determine sequence number: first element inherits from the parent call, the rest is 1
				seq := 1
				if i == 0 {
					seq = sequenceNumber
				}
				// Parse the line and increment the line index
				nameOk, err := ParseLine(inputLines[*lineIndex], targetValue, targetStructAnnotation, seq, config)
				*lineIndex++
				if err != nil {
					return err
				}
				// If there is a type name mismatch but the target is optional it can be skipped, otherwise it's an error
				if !nameOk {
					if _, exists := targetStructAnnotation.Attributes[constants.AttributeOptional]; exists {
						err = nil
						*lineIndex--
						continue
					} else {
						return fmt.Errorf("%w @ln %d", errmsg.ErrStructureParsingLineTypeNameMismatch, *lineIndex)
					}
				}
			}
		}
	}
	// Return nil if everything went well
	return nil
}
