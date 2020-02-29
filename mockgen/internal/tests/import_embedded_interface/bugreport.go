// Copyright 2020 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
//go:generate mockgen -destination bugreport_mock.go -package bugreport -source=bugreport.go

package bugreport

import (
	"log"

	"github.com/golang/mock/mockgen/internal/tests/import_embedded_interface/ersatz"
	"github.com/golang/mock/mockgen/internal/tests/import_embedded_interface/faux"
)

// Source is an interface w/ an embedded foreign interface
type Source interface {
	ersatz.Embedded
	faux.Foreign
}

func CallForeignMethod(s Source) {
	log.Println(s.Ersatz())
	log.Println(s.OtherErsatz())
}
