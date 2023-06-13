# Changelog

## [1.6.3](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.6.2...v1.6.3) (2023-06-13)


### Bug Fixes

* field possible unknown for RequireIfAttributeIsOneOf ([6a4935f](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/6a4935f8548b483c52a6bd91f13c85b0a4161d73))
* missing path expression in NullIfAttributeIsOneOf ([3b6668d](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/3b6668d66e956b989386a0cd29ff7c1d5f2b3316))


### Miscellaneous

* **deps:** bump github.com/hashicorp/terraform-plugin-framework ([#62](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/62)) ([cc93eca](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/cc93ecabbe6b55bdaed3d55e2af1725a3cf4e433))

## [1.6.2](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.6.1...v1.6.2) (2023-06-09)


### Bug Fixes

* bad func used for nullattributeisoneof and fix description ([#60](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/60)) ([0411638](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/04116382b39e8ab72fad520a17a1940b17e999b7))

## [1.6.1](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.6.0...v1.6.1) (2023-06-09)


### Bug Fixes

* attributeIsDivisibleByAnInteger bad check if value IsNull and IsUnknown ([#58](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/58)) ([7b8fd83](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/7b8fd8311024606d4f76361bf31131feaff74920))

## [1.6.0](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.5.2...v1.6.0) (2023-06-09)


### Features

* add validator `null_if_attribute_is_one_of` ([#52](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/52)) ([ada0563](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/ada056330e16d045868d187abd703e7c86923911))


### Bug Fixes

* pathExpression missing in attribute NullIfAttributeIsOneOf validator ([#57](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/57)) ([6b0e629](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/6b0e6291635435fcc84715b3a4bd4533621bdf0f))

## [1.5.2](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.5.1...v1.5.2) (2023-06-08)


### Bug Fixes

* pathExpression for RequireIfAttributeIsOneOf validator ([#53](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/53)) ([2058e4c](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/2058e4c01947a90d0d923968a814aeca4532acbf))
* pathExpression missing in attributeIsDivisibleByAnInteger validator ([e99edac](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/e99edac0042a7198ee0d30fb1a38c70b307dc006))

## [1.5.1](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.5.0...v1.5.1) (2023-04-25)


### Bug Fixes

* RequireIfAttributeIsOneOf error message and tests ([f857608](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/f857608557d4b7999ec365eba8744129cd870b8e))

## [1.5.0](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.4.0...v1.5.0) (2023-04-14)


### Features

* add string and int64 validator `OneOfWithDescription` ([23b6925](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/23b692537a9220243c2d91e16722755b7b25df25))

## [1.4.0](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.3.1...v1.4.0) (2023-04-07)


### Features

* add int64 validator `AttributeIsDevidedByAnInteger` ([#46](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/46)) ([fa6fcbe](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/fa6fcbee7f6152f295e03f231353baa0957dd3f2))
* add int64 validator `ZeroRemainder` ([#47](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/47)) ([d7c6d19](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/d7c6d1961a02eeb8639b623ef9bf5ad352baa0fa))


### Miscellaneous

* CODE_OF_CONDUCT ([1484d02](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/1484d02b4987051f86bed62645789ee99fbdd509))
* LICENCE ([1d4c81a](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/1d4c81ac0a8b5f19b5b77cbb9cec9015c6ded3d2))

## [1.3.1](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/compare/v1.3.0...v1.3.1) (2023-03-31)


### Bug Fixes

* error message and docs for `RequireIfAttributeIsOneOf` ([#42](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/issues/42)) ([b1b3096](https://github.com/FrangipaneTeam/terraform-plugin-framework-validators/commit/b1b30966e6fb3fe51177af27dd7994b2669381ab))

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
