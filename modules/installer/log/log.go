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

package log

import "github.com/sirupsen/logrus"

type Logger interface {
	Infof(string, ...interface{})
	Debugf(string, ...interface{})
	Errorf(string, ...interface{})
	Fatalf(string, ...interface{})
}

func NewLogger(debug bool) Logger {
	l := logrus.New()
	if debug {
		l.SetLevel(logrus.DebugLevel)
	}

	timeFormatter := new(logrus.TextFormatter)
	timeFormatter.ForceColors = true
	timeFormatter.FullTimestamp = true
	timeFormatter.TimestampFormat = "2006-01-02 15:04:05"
	l.SetFormatter(timeFormatter)

	return l
}
