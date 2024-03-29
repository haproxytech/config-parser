// Code generated by go generate; DO NOT EDIT.
/*
Copyright 2019 HAProxy Technologies

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

package integration_test

const userlist_groupG1userstigerscott = `
userlist test
  group G1 users tiger,scott
`
const userlist_groupG1 = `
userlist test
  group G1
`
const userlist_usertigerpassword6k6y3oePJlKBxxH = `
userlist test
  user tiger password $6$k6y3o.eP$JlKBx(...)xHSwRv6J.C0/D7cV91 groups G1
`
const userlist_userpandainsecurepasswordelgatog = `
userlist test
  user panda insecure-password elgato groups G1,G2
`
const userlist_userbearinsecurepasswordhellogro = `
userlist test
  user bear insecure-password hello groups G2
`
const userlist_userplatipusinsecurepasswordsalu = `
userlist test
  user platipus insecure-password saludos
`
