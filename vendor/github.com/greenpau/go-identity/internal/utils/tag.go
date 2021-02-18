package utils

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"reflect"
	"strings"
)

func GetTagCompliance(resource interface{}) ([]string, bool) {
	result := true
	output := []string{}

	resourceInstance := reflect.TypeOf(resource).Elem()
	//resourceType := fmt.Sprintf("%s", resourceInstance.Name())
	resourceKind := fmt.Sprintf("%s", resourceInstance.Kind())

	if resourceKind != "struct" {
		output = append(output, fmt.Sprintf(
			"FAIL: %s resource kind is unsupported",
			resourceKind,
		))
		return output, false
	}

	suggestedStructChanges := []string{}

	requiredTags := []string{"json", "xml", "yaml"}
	for i := 0; i < resourceInstance.NumField(); i++ {
		resourceField := resourceInstance.Field(i)
		expTagValue := convertFieldToTag(resourceField.Name)
		expTagValue = expTagValue + ",omitempty"
		var lastTag bool
		for j, tagName := range requiredTags {
			if len(requiredTags)-1 == j {
				lastTag = true
			}

			tagValue := resourceField.Tag.Get(tagName)

			if tagValue == "-" {
				continue
			}
			if tagValue == "" {
				result = false
				output = append(output, fmt.Sprintf(
					"FAIL: %s tag not found in %s.%s (%v)",
					tagName,
					//resourceType,
					resourceInstance.Name(),
					resourceField.Name,
					resourceField.Type,
				))
				if lastTag {
					tags := makeTags(requiredTags, expTagValue)
					resType := fmt.Sprintf("%v", resourceField.Type)
					resType = strings.Replace(resType, "identity.", "", -1)
					suggestedStructChanges = append(suggestedStructChanges, fmt.Sprintf(
						"%s %s %s", resourceField.Name, resType, tags,
					))
				}
				continue
			}
			//if strings.Contains(tagValue, ",omitempty") {
			//	tagValue = strings.Replace(tagValue, ",omitempty", "", -1)
			//}
			if tagValue != expTagValue {
				result = false
				output = append(output, fmt.Sprintf(
					"FAIL: %s tag mismatch found in %s.%s (%v): %s (actual) vs. %s (expected)",
					tagName,
					//resourceType,
					resourceInstance.Name(),
					resourceField.Name,
					resourceField.Type,
					tagValue,
					expTagValue,
				))
				continue

			}
		}
	}

	if len(suggestedStructChanges) > 0 {
		output = append(output, fmt.Sprintf(
			"FAIL: suggested struct changes to %s:\n%s",
			resourceInstance.Name(),
			strings.Join(suggestedStructChanges, "\n"),
		))
	}

	return output, result
}

func convertFieldToTag(s string) string {
	s = strcase.ToSnake(s)
	s = strings.ReplaceAll(s, "_md_5", "_md5")
	s = strings.ReplaceAll(s, "open_ssh", "openssh")
	return s
}

func makeTags(tags []string, s string) string {
	var b strings.Builder
	b.WriteRune('`')
	tagOutput := []string{}
	for _, tag := range tags {
		tagOutput = append(tagOutput, tag+":\""+s+"\"")
	}
	b.WriteString(strings.Join(tagOutput, " "))
	b.WriteRune('`')
	return b.String()
}
