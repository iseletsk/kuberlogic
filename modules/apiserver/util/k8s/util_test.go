/*
 * CloudLinux Software Inc 2019-2021 All Rights Reserved
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package k8s

import "testing"

func TestMapToStrSelector(t *testing.T) {
	k8sSelector := map[string]string{"app": "test", "env": "testing"}
	k8sSelectorExpected := "app=test,env=testing"
	k8sSelectorActual := MapToStrSelector(k8sSelector)

	if k8sSelectorActual != k8sSelectorExpected {
		t.Errorf("failed, got %s, want %s", k8sSelectorActual, k8sSelectorExpected)
	}
}
