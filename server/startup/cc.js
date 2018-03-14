/**
 * @license
 * Code City: Startup code.
 *
 * Copyright 2017 Google Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/**
 * @fileoverview Initialisation code to set up CC-specific extensions.
 * @author cpcallen@google.com (Christopher Allen)
 */

// Threads API; parts are roughly conformant with HTML Living
// Standard, plus our local extensions:

var suspend = new 'suspend';
var setTimeout = new 'setTimeout';
var clearTimeout = new 'clearTimeout';

// Namespace for CodeCity-specific extensions:
var CC = {};

// Networking functions.
CC.connectionListen = new 'CC.connectionListen';
CC.connectionUnlisten = new 'CC.connectionUnlisten';
CC.connectionWrite = new 'CC.connectionWrite';
CC.connectionClose = new 'CC.connectionClose';