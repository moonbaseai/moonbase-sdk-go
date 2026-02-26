# Changelog

## 0.1.0-alpha.9 (2026-02-26)

Full Changelog: [v0.1.0-alpha.8...v0.1.0-alpha.9](https://github.com/moonbaseai/moonbase-sdk-go/compare/v0.1.0-alpha.8...v0.1.0-alpha.9)

### Bug Fixes

* **types:** correct string fields to enums in activity/call/meeting ([2da8b0d](https://github.com/moonbaseai/moonbase-sdk-go/commit/2da8b0d7838ff36ab3b4ef7e6d9f3c8fa7814617))


### Chores

* **internal:** move custom custom `json` tags to `api` ([7a2c48c](https://github.com/moonbaseai/moonbase-sdk-go/commit/7a2c48c427963b3f46501e97ae0140bf670e8ee2))

## 0.1.0-alpha.8 (2026-02-23)

Full Changelog: [v0.1.0-alpha.7...v0.1.0-alpha.8](https://github.com/moonbaseai/moonbase-sdk-go/compare/v0.1.0-alpha.7...v0.1.0-alpha.8)

### Features

* **api:** add search to collections/items, add Search resource, remove Items resource ([88edee3](https://github.com/moonbaseai/moonbase-sdk-go/commit/88edee3d341e740bae3da3bffbae725d013b6827))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([777e93e](https://github.com/moonbaseai/moonbase-sdk-go/commit/777e93e6a0c318a2a5f92ae9022e14ad2d131119))
* **client:** use correct format specifier for header serialization ([9e226bf](https://github.com/moonbaseai/moonbase-sdk-go/commit/9e226bf1b8b34842a057dce6755770e984d154ce))
* **encoder:** correctly serialize NullStruct ([7e84631](https://github.com/moonbaseai/moonbase-sdk-go/commit/7e846310aca2789523f94d28d76cc6c326116347))


### Chores

* update mock server docs ([3f79a3b](https://github.com/moonbaseai/moonbase-sdk-go/commit/3f79a3b26a234e40e0e7778877d651d67d1a3894))

## 0.1.0-alpha.7 (2026-02-09)

Full Changelog: [v0.1.0-alpha.6...v0.1.0-alpha.7](https://github.com/moonbaseai/moonbase-sdk-go/compare/v0.1.0-alpha.6...v0.1.0-alpha.7)

### Features

* **api:** update api ([b0af728](https://github.com/moonbaseai/moonbase-sdk-go/commit/b0af7284002199e1c5dc268404e7e2824bba7dff))

## 0.1.0-alpha.6 (2026-02-02)

Full Changelog: [v0.1.0-alpha.5...v0.1.0-alpha.6](https://github.com/moonbaseai/moonbase-sdk-go/compare/v0.1.0-alpha.5...v0.1.0-alpha.6)

### Features

* **api:** manual updates ([c75dc0e](https://github.com/moonbaseai/moonbase-sdk-go/commit/c75dc0ed748f394b3db62c2cad6228f071a535b4))
* **api:** update api ([3e2bf65](https://github.com/moonbaseai/moonbase-sdk-go/commit/3e2bf658d23793ce67931689658b13a775e0169d))
* **api:** update api ([7e9bd90](https://github.com/moonbaseai/moonbase-sdk-go/commit/7e9bd90a2aa1c53c87f4e0bc44331572402a9e6f))
* **api:** update api ([0fd2c56](https://github.com/moonbaseai/moonbase-sdk-go/commit/0fd2c56e2455abd875d070c53bbb238175548b9b))
* **api:** update api ([64823fd](https://github.com/moonbaseai/moonbase-sdk-go/commit/64823fdf34e4f3dab18a32384f4c4c90351ddd2a))
* **api:** update api ([9a08348](https://github.com/moonbaseai/moonbase-sdk-go/commit/9a0834831afbab009cedfe9af74203c0072b7f3d))
* **api:** update api ([cbb8c67](https://github.com/moonbaseai/moonbase-sdk-go/commit/cbb8c6732da4e13886ef94be2784557d0e56f5f1))
* **api:** update api ([735ae8d](https://github.com/moonbaseai/moonbase-sdk-go/commit/735ae8de247a868d88626bbb08f051b816c391cb))
* **api:** update api ([9475b65](https://github.com/moonbaseai/moonbase-sdk-go/commit/9475b65b888ef9c0c736a39945b4c65808ddd846))
* **api:** update api ([b64801a](https://github.com/moonbaseai/moonbase-sdk-go/commit/b64801a74cd85d173bd409af7ca9c53f26096f7a))
* **client:** add a convenient param.SetJSON helper ([4f33e36](https://github.com/moonbaseai/moonbase-sdk-go/commit/4f33e36606b3d3a7f96abe651f9a51404e1e1985))
* **encoder:** support bracket encoding form-data object members ([f98da4f](https://github.com/moonbaseai/moonbase-sdk-go/commit/f98da4f84e56f221c54a08a43d2cd086f9796c95))


### Bug Fixes

* **client:** correctly specify Accept header with */* instead of empty ([ca87c33](https://github.com/moonbaseai/moonbase-sdk-go/commit/ca87c3301fb481d3a7b6721aedfe48867fc776b2))
* **docs:** add missing pointer prefix to api.md return types ([aff0f15](https://github.com/moonbaseai/moonbase-sdk-go/commit/aff0f157db1da4723ed4b14543c1da1ce35d6572))
* **docs:** fix mcp installation instructions for remote servers ([f83bdff](https://github.com/moonbaseai/moonbase-sdk-go/commit/f83bdff7d1fc87bed3fe20f156c583675d46eaef))
* **mcp:** correct code tool API endpoint ([040cc7c](https://github.com/moonbaseai/moonbase-sdk-go/commit/040cc7c26f8d37cd19631a25f02d7e755964c89d))
* rename param to avoid collision ([5fd2db8](https://github.com/moonbaseai/moonbase-sdk-go/commit/5fd2db8acb31b38a46e894912f929339cfbb02a5))


### Chores

* add float64 to valid types for RegisterFieldValidator ([db9e501](https://github.com/moonbaseai/moonbase-sdk-go/commit/db9e5018dca229002b7c516ee0a254172f845220))
* elide duplicate aliases ([b9acba3](https://github.com/moonbaseai/moonbase-sdk-go/commit/b9acba3bc29d7a5cd43c8179775d5cf4d4f8423d))
* **internal:** codegen related update ([1145815](https://github.com/moonbaseai/moonbase-sdk-go/commit/11458159d4689299e355bf90101620ffb3b893fa))
* **internal:** codegen related update ([16d9bfc](https://github.com/moonbaseai/moonbase-sdk-go/commit/16d9bfc9513037c2d8fae62de2762f56994820c7))
* **internal:** update `actions/checkout` version ([7489cd1](https://github.com/moonbaseai/moonbase-sdk-go/commit/7489cd17363dfa9633b2d97077c13fb34162d6e4))


### Documentation

* **api:** update meeting params ([c58852c](https://github.com/moonbaseai/moonbase-sdk-go/commit/c58852ce688e21c74f34a9ec5b82bbcf6f2276ef))
* prominently feature MCP server setup in root SDK readmes ([7fc4a2a](https://github.com/moonbaseai/moonbase-sdk-go/commit/7fc4a2aecb0ba7d8552a42b4f4b405497b7f03b3))

## 0.1.0-alpha.5 (2025-11-13)

Full Changelog: [v0.1.0-alpha.4...v0.1.0-alpha.5](https://github.com/moonbaseai/moonbase-sdk-go/compare/v0.1.0-alpha.4...v0.1.0-alpha.5)

### Features

* Add PATCH /v0/meetings/{id} ([a29d08f](https://github.com/moonbaseai/moonbase-sdk-go/commit/a29d08f1740b2766d842010c68b696356afb9238))
* **api:** manual updates ([a499647](https://github.com/moonbaseai/moonbase-sdk-go/commit/a499647162c655a3acfcb0348509500b536a6d58))
* **api:** update api ([3bcd3bf](https://github.com/moonbaseai/moonbase-sdk-go/commit/3bcd3bf0bdb808bbcdac8af33a2331aa0e69ac73))
* **client:** handle recursive schemas properly ([5e4732a](https://github.com/moonbaseai/moonbase-sdk-go/commit/5e4732a777f111e547a26b6a3afaf866f23a8273))


### Bug Fixes

* bugfix for setting JSON keys with special characters ([b8b63ec](https://github.com/moonbaseai/moonbase-sdk-go/commit/b8b63ec36047b22190d913a6995cef7b0c5e2b5b))
* use slices.Concat instead of sometimes modifying r.Options ([4ea87ce](https://github.com/moonbaseai/moonbase-sdk-go/commit/4ea87ce0342fad87654de8a05094d6c247398097))


### Chores

* bump gjson version ([3cf4c58](https://github.com/moonbaseai/moonbase-sdk-go/commit/3cf4c583efcba2a05346bd3d2347392647155479))
* bump minimum go version to 1.22 ([6d1a4d9](https://github.com/moonbaseai/moonbase-sdk-go/commit/6d1a4d9a1366bf323651433dc1cf9d8b4a782223))
* do not install brew dependencies in ./scripts/bootstrap by default ([f9571b1](https://github.com/moonbaseai/moonbase-sdk-go/commit/f9571b17533ed06bdbd5592ad177028f9a08cc0c))
* **internal:** grammar fix (it's -&gt; its) ([17d7fff](https://github.com/moonbaseai/moonbase-sdk-go/commit/17d7fff00a96cdbaba0810a7d963688e9cd97dcf))
* update more docs for 1.22 ([2eb338c](https://github.com/moonbaseai/moonbase-sdk-go/commit/2eb338c5fb52b2f9b7ba6be164365ea68e6a931e))

## 0.1.0-alpha.4 (2025-09-12)

Full Changelog: [v0.1.0-alpha.3...v0.1.0-alpha.4](https://github.com/moonbaseai/moonbase-sdk-go/compare/v0.1.0-alpha.3...v0.1.0-alpha.4)

### Features

* Improve examples of API errors ([f9cf183](https://github.com/moonbaseai/moonbase-sdk-go/commit/f9cf1833e99d8324372f4bc38de0b5da13e48472))


### Documentation

* improve webhook endpoints examples ([5f7104f](https://github.com/moonbaseai/moonbase-sdk-go/commit/5f7104f06a0e55d0d4b99efbf0fba6ab86215b84))


### Refactors

* rename ChoiceFieldOption label to name for consistency ([b52ec31](https://github.com/moonbaseai/moonbase-sdk-go/commit/b52ec311ec1209339371db5bef24f2630bb21384))

## 0.1.0-alpha.3 (2025-09-09)

Full Changelog: [v0.1.0-alpha.2...v0.1.0-alpha.3](https://github.com/moonbaseai/moonbase-sdk-go/compare/v0.1.0-alpha.2...v0.1.0-alpha.3)

### Features

* **api:** add upsert endpoint for Calls ([ed9cebf](https://github.com/moonbaseai/moonbase-sdk-go/commit/ed9cebf19eb9b608cfcdc223a573837bb9e3e76a))
* **api:** example updates ([555a675](https://github.com/moonbaseai/moonbase-sdk-go/commit/555a675f764f6191b126965aa37638aab3a14275))
* **api:** manual updates ([0a103e8](https://github.com/moonbaseai/moonbase-sdk-go/commit/0a103e81c6c2c488a6aa07e760552e4b506bb14e))
* **api:** manual updates ([c0e6c48](https://github.com/moonbaseai/moonbase-sdk-go/commit/c0e6c48a59497dace54bdbac9773eb50cc766047))
* **api:** update api ([7c2e6c6](https://github.com/moonbaseai/moonbase-sdk-go/commit/7c2e6c609e2ff58ffa8cd2804af019c09d1f6a13))
* **api:** update api ([3955136](https://github.com/moonbaseai/moonbase-sdk-go/commit/39551364aa514a1659a9b0f71c04f8389961dacd))
* **api:** update api ([0e2a55f](https://github.com/moonbaseai/moonbase-sdk-go/commit/0e2a55fb9923bab332d1e4d77df0fa870dff2743))
* **api:** update api ([4bb3353](https://github.com/moonbaseai/moonbase-sdk-go/commit/4bb3353d00228b23186130a1641c4fdfcac10698))
* **api:** update api ([3935b24](https://github.com/moonbaseai/moonbase-sdk-go/commit/3935b24b94b1618b27ede37e1e208d574a3f69bb))
* **api:** update examples ([68f3fbf](https://github.com/moonbaseai/moonbase-sdk-go/commit/68f3fbffc1fa9d344bcd1140c5c46e60801571c7))
* **client:** support optional json html escaping ([dd43aa6](https://github.com/moonbaseai/moonbase-sdk-go/commit/dd43aa60d6b11a4e73b825d9d5a10c7127c42625))


### Bug Fixes

* **client:** process custom base url ahead of time ([c7920a2](https://github.com/moonbaseai/moonbase-sdk-go/commit/c7920a250cd506e113853b5b426ade26ab4db153))
* close body before retrying ([2a573bf](https://github.com/moonbaseai/moonbase-sdk-go/commit/2a573bf19be18079393c6b2f6630579afaea1644))
* **internal:** unmarshal correctly when there are multiple discriminators ([0fa643d](https://github.com/moonbaseai/moonbase-sdk-go/commit/0fa643db59749a0739bd88469f31a0945229e9d9))


### Chores

* **internal:** codegen related update ([273f6f4](https://github.com/moonbaseai/moonbase-sdk-go/commit/273f6f4fa7383c69a694213e20f083ed7bedb65e))
* **internal:** codegen related update ([0aec83a](https://github.com/moonbaseai/moonbase-sdk-go/commit/0aec83ad62d58c0c895761db318b99cd5cd1a93d))
* **internal:** update comment in script ([435366b](https://github.com/moonbaseai/moonbase-sdk-go/commit/435366b558e67de00049c6cc91d9546ee59930dc))
* update @stainless-api/prism-cli to v5.15.0 ([cc3a342](https://github.com/moonbaseai/moonbase-sdk-go/commit/cc3a34228df9a371ba34a70d91a593f1c3e166c8))

## 0.1.0-alpha.2 (2025-07-20)

Full Changelog: [v0.1.0-alpha.1...v0.1.0-alpha.2](https://github.com/moonbaseai/moonbase-sdk-go/compare/v0.1.0-alpha.1...v0.1.0-alpha.2)

### Bug Fixes

* pagination ([4186ff3](https://github.com/moonbaseai/moonbase-sdk-go/commit/4186ff349decddb0891b3c19530366bc30adba94))

## 0.1.0-alpha.1 (2025-07-18)

Full Changelog: [v0.0.1-alpha.0...v0.1.0-alpha.1](https://github.com/moonbaseai/moonbase-sdk-go/compare/v0.0.1-alpha.0...v0.1.0-alpha.1)

### Features

* **api:** update api ([239b955](https://github.com/moonbaseai/moonbase-sdk-go/commit/239b9555be52b88f8800d7323d88f1e49cb5f2b0))
* **api:** update api ([e200325](https://github.com/moonbaseai/moonbase-sdk-go/commit/e2003255a586dab1dec331fe5b0fb0a64e465274))
* **api:** update via SDK Studio ([5cccd87](https://github.com/moonbaseai/moonbase-sdk-go/commit/5cccd87a0cfc8da4a2c521902c213e8443aa7445))
* **api:** update via SDK Studio ([020a150](https://github.com/moonbaseai/moonbase-sdk-go/commit/020a150c5473bfea73b81b2364f2c7540290d82e))
* **api:** update via SDK Studio ([490cedc](https://github.com/moonbaseai/moonbase-sdk-go/commit/490cedc52f47a56dc75bf2b953bba4b74b3f8864))
* **api:** update via SDK Studio ([7903747](https://github.com/moonbaseai/moonbase-sdk-go/commit/79037473add553870a1f8169a6fb4d8b17ff4a88))
* **api:** update via SDK Studio ([38d7a8b](https://github.com/moonbaseai/moonbase-sdk-go/commit/38d7a8bc82ba9f68352af24f196e4ec2f2abb7d0))
* **api:** update via SDK Studio ([3a66b48](https://github.com/moonbaseai/moonbase-sdk-go/commit/3a66b48faf2291ce0504acd44ecb1f2b25664ae6))
* **api:** update via SDK Studio ([fa9be42](https://github.com/moonbaseai/moonbase-sdk-go/commit/fa9be423c9ebbb87582c037304b64addacd51efb))
* **api:** update via SDK Studio ([ffa23e5](https://github.com/moonbaseai/moonbase-sdk-go/commit/ffa23e5a76d0fb238b56dedf93b2d18349c38e73))
* **api:** update via SDK Studio ([1dc3346](https://github.com/moonbaseai/moonbase-sdk-go/commit/1dc3346b6ea9a09663000c64e6e541063069338d))
* **api:** update via SDK Studio ([478b226](https://github.com/moonbaseai/moonbase-sdk-go/commit/478b22600e5b50d2da831f8a5e111c7e9e296acd))
* **api:** update via SDK Studio ([b6f11fd](https://github.com/moonbaseai/moonbase-sdk-go/commit/b6f11fdc66d8db8c2302e945cfaaf0b5bc467e19))
* **api:** update via SDK Studio ([98adad5](https://github.com/moonbaseai/moonbase-sdk-go/commit/98adad5161c501c238b87389825c94add6359e64))
* **api:** update via SDK Studio ([3bc9a71](https://github.com/moonbaseai/moonbase-sdk-go/commit/3bc9a71a8a1d512e8f6ffb356a3eff0dbb34c970))
* **api:** update via SDK Studio ([707155c](https://github.com/moonbaseai/moonbase-sdk-go/commit/707155c8ee8ca961468f0316f62a4c22d904f2cc))
* **api:** update via SDK Studio ([350a89e](https://github.com/moonbaseai/moonbase-sdk-go/commit/350a89e71398e4f5ca46ba64eadff074c864d480))
* **api:** update via SDK Studio ([d0fc1b6](https://github.com/moonbaseai/moonbase-sdk-go/commit/d0fc1b630ef79d3358fb86e57f6b2d81f24021af))
* **api:** update via SDK Studio ([4cc5570](https://github.com/moonbaseai/moonbase-sdk-go/commit/4cc5570467a09fb30a99afa16a0383f0a78dd65b))
* **api:** update via SDK Studio ([43f53a6](https://github.com/moonbaseai/moonbase-sdk-go/commit/43f53a69ee7c1a352dd48498802df7bc46557e44))
* **api:** update via SDK Studio ([f8390e6](https://github.com/moonbaseai/moonbase-sdk-go/commit/f8390e678a9c0fd40606bff79e2cacd827c1f1d8))
* **api:** update via SDK Studio ([2c272cf](https://github.com/moonbaseai/moonbase-sdk-go/commit/2c272cf4ecc790ba20dd5807be2ffa3f22bb08cb))


### Bug Fixes

* Program and ProgramTemplate circular reference ([1b91086](https://github.com/moonbaseai/moonbase-sdk-go/commit/1b910863879492e1627f7d81e41bd7cb3e341317))


### Chores

* Update license ([1f19954](https://github.com/moonbaseai/moonbase-sdk-go/commit/1f1995472edff649a1f00ff180fe2f9017c7a2f3))
* update SDK settings ([d597c3c](https://github.com/moonbaseai/moonbase-sdk-go/commit/d597c3c7fe4e52afff5503ca026e343f05a4086f))
