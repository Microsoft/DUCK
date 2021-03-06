// Data Use Statement Compliance Checker (DUCK)
// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

/**
 * Defines application modules
 */
angular.module('duck.core', []);
angular.module("duck.component", []);
angular.module("duck.i18n", []);
angular.module("duck.gateway", []);
angular.module("duck.assumption", []);
angular.module("duck.main", ["ui.router", "ngAnimate"]);
angular.module("duck.editor", ["MassAutoComplete", "puElasticInput", "ngSanitize", "ui.router", "ngFileSaver"]);
angular.module("duck.signin", ["ui.router"]);
angular.module("duck.home", ["ui.router"]);

