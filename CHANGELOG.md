# Changelog

## [1.3.0](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.2.1...v1.3.0) (2023-03-31)


### Features

* add new validator `require_if_attribute_is_one_of` ([b685181](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/b685181ca9f1f35bf14dbd7851f9a69bbe6040e8))

## [1.2.1](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.2.0...v1.2.1) (2023-03-29)


### Bug Fixes

* Missing documentation for `IsMacAddress` string validator ([#35](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/35)) ([18feb10](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/18feb1039d28915516aa62a1ba0dbf87b9f0bbea))


### Documentation

* merge doc to one projet ([08b6356](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/08b635652f5d93f962f73b1af95780342c0a8ce1))


### Miscellaneous

* **deps:** bump github.com/hashicorp/terraform-plugin-framework ([4470618](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/447061860b918c64b78ac4f358b9e764cc303289))
* **docs:** Add docs ValueStringsAre in list/set/map ([10472f7](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/10472f785d1a34c56c83d92936c8cbe6e6752889))
* **docs:** Add Favicon and branding orange ([d6a7005](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/d6a7005a325fe22a6b34f254b7cdd645d59233cf))
* remove list IsURN validator ([10472f7](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/10472f785d1a34c56c83d92936c8cbe6e6752889))

## [1.2.0](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.1.1...v1.2.0) (2023-03-20)


### Features

* Add `IsMacAddress` String validator ([4a82bc8](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/4a82bc81df65ea7d5a7fae1da7af4217405751f9))

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
