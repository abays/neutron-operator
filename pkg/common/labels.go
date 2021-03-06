/*
Copyright 2020 Red Hat

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package common

// GetLabels - get labels to be set on objects created by controller
func GetLabels(name string, appLabel string) map[string]string {
	return map[string]string{
		"owner": "neutron-operator",
		"cr":    name,
		"app":   appLabel,
	}
}

// InitLabelMap - Inititialise a label map to an empty map if it is nil.
func InitLabelMap(m *map[string]string) {
	if *m == nil {
		*m = make(map[string]string)
	}
}
