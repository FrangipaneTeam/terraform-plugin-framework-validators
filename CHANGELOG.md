# Changelog

## [1.1.1](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.1.0...v1.1.1) (2023-03-20)


### Miscellaneous

* **docs:** Add docs ValueStringsAre in list/set/map ([033ee2e](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/033ee2e0abce511baca5697fc9c7b4a4bdabc6c4))
* remove list IsURN validator ([033ee2e](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/033ee2e0abce511baca5697fc9c7b4a4bdabc6c4))

## [1.1.0](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.0.0...v1.1.0) (2023-03-17)


### Features

* add `IsURN` list validator ([#20](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/20)) ([9fbe5d2](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/9fbe5d2121f3f215436c7b48ebdd1c2a765abb12))


### Bug Fixes

* better diagnostic for listvalidator isURN ([#26](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/26)) ([49ccf56](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/49ccf5621a36dace5fbd422a061e17d7d60f93d5))


### Documentation

* add links to planmodifier and to homepage ([d5ce111](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/d5ce111056f2e25a28b12b7cab717dc70631fdd6))
* add released on each validators ([3ffd70c](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/3ffd70c69aa8b53f4c99f6380cb4b56241e1fbd5))

## 1.0.0 (2023-03-13)


### Features

* add `IsValidURN` ([096d505](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/096d50591bb539d8c659aafa31e14427f3c43212))
* add `IsValidUUID` string validator ([99efea0](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/99efea0db4fad76c66b2c81d3d498f329eba8824))
* add `Not` bool validator ([ed07f42](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/ed07f42d28177f49659ab8e3f05c97a48aaea3cd))
* add `Not` int64 validator ([7ff00f0](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/7ff00f026f6a849547e35ec7c4a6f8a2b9adc747))
* add `Not` list validator ([ab8bcd8](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/ab8bcd850dc1bc75bed749b87d66c74ea2888b89))
* add `Not` map validator ([b779fe6](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/b779fe6cd92723c4b200192303eba3e6794f4960))
* add `Not` set validator ([3581f9d](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/3581f9df692c7e9a9ab3ae86805f6d10e8cc0a59))
* add `Not` string validator ([18558de](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/18558deb8d09a1d983a553d4a6a7b182847b4fa0))
* **netmask:** add validator `IsValidNetmask` ([fc2452c](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/fc2452c8db78db1673852d0a956b7981326c5d37))
* valid ip with net.ParseIP ([dc064f6](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/dc064f6846d6c7b5c995f40e519b6a4d04871f6c))


### Documentation

* add usage documentation ([86a8219](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/86a8219b2eaed848c7a66a223c24b3bfe49e5e21))


### Miscellaneous

* add CI, Docs, ReleasePlease, Markdownlint ([a4517d0](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/a4517d0829737441f39e2cb21abf7e0c10c9fec6))
