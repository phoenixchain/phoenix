package systemcontracts

import (
	"bytes"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

const (
	// genesis contracts
	ValidatorHubContract     = "0x0000000000000000000000000000000000001006"
	ValidatorFactoryContract = "0x0000000000000000000000000000000000001008"
)

var (
	SystemContracts = map[common.Address]bool{
		common.HexToAddress(ValidatorHubContract):     true,
		common.HexToAddress(ValidatorFactoryContract): true,
	}

	SystemContractAddress = map[common.Address]bool{
		common.HexToAddress(ValidatorHubContract):                         true,
		common.HexToAddress(ValidatorFactoryContract):                     true,
		common.HexToAddress("0xA020c0a38953C7E8Dafea49e8a3c4607130DDa66"): true,
		common.HexToAddress("0xc07B539C53A07CEa4E26C6d7c841cE200E0aBc2C"): true,
		common.HexToAddress("0x02fD225919FB8Adc7275273B280a9B3b01e13fce"): true,
		common.HexToAddress("0xF2Aec47A2AB627Fa32368b04E33395908e9D3349"): true,
		common.HexToAddress("0xf3c0fe6cacC16Bb717759701d8BfD60566aa7700"): true,
		common.HexToAddress("0x07763c8479C6BD80d76e1b76Aec6e2C8225fA3Aa"): true,
		common.HexToAddress("0x95819AA15A3D7E076807673172cd62F50DcbAC8F"): true,
		common.HexToAddress("0x1bd1dB8034AAc8966E889b0Bdc31f7D89fB3a4a9"): true,
		common.HexToAddress("0x9b4a11a6c67b167e856503344c469D0dD6385917"): true,
		common.HexToAddress("0x9931E8eEb56B8948E671ec5E7ae3bd6f1D887599"): true,
		common.HexToAddress("0x9AAdb82b6E04A90B16b22bC84aFFa96bD4108b1d"): true,
		common.HexToAddress("0xcADBE233f28bD95b9FAa5cf0fcB5662B78af141F"): true,
		common.HexToAddress("0x73EA56cD46304a9fD6801a8cb87a49828F2Ea5EA"): true,
		common.HexToAddress("0xC867E26b913F8287A5C35f16b8448EcFb195D07C"): true,
		common.HexToAddress("0xAC521550074D2c12430d7C7fbdEE435ef0Cd901A"): true,
		common.HexToAddress("0xb592D778511D2F336c17286cD1F64F6e57fC7eF5"): true,
		common.HexToAddress("0x3f12c701Cf9b2498594FDb2000E2000c38fDA24f"): true,
		common.HexToAddress("0x4cE9a0f86021b0a65CDA6Be90363CD69eE2Cdd6E"): true,
		common.HexToAddress("0x21dAb718b89cE13a5F2AC7c861E7F4f64eC25B6C"): true,
		common.HexToAddress("0x05397a58357d57629AE5305828268f4e814b3565"): true,
		common.HexToAddress("0x2fE3A306b8BA6ffdf0252C5d2003b16839b199a8"): true,
		common.HexToAddress("0xe15E988b1D89b021C4E3176fb376e3E6c18C471f"): true,
		common.HexToAddress("0x459468063C7819A5Dd5354F64E6f89FB3F163F37"): true,
		common.HexToAddress("0x126ca6A66319c17930b750504DCF18ADB41C5d13"): true,
		common.HexToAddress("0x32Aa65870388F4f0bBe95D2F05C40b25d8a0f285"): true,
		common.HexToAddress("0x17018dB830D6180bD526458704A866FE5582A750"): true,
		common.HexToAddress("0x259Bb7e3172FCB91C7bd9329B7898bF59C0B0351"): true,
		common.HexToAddress("0xDDBa46900859c97d53623697b40588E756be9016"): true,
		common.HexToAddress("0x924CEd3131E2fB3C07cb589870D324227359f943"): true,
		common.HexToAddress("0xe56dF61098eca58f23Cf1622cA7a53F352A9318c"): true,
		common.HexToAddress("0x3F44cfAa776eeD03f43a783A8D244e5CB4566c79"): true,
		common.HexToAddress("0x9a3F8d9D230281eb9f238436c396c71B8AD1C08E"): true,
		common.HexToAddress("0xb52FacbFbE215591D8E9dF495a9cf202B3e422d0"): true,
		common.HexToAddress("0xd69a3d2D4eA50c877E5568A589967A73b81c50dC"): true,
		common.HexToAddress("0x94905830148c54E71a3291e2379307C7b204fAa8"): true,
		common.HexToAddress("0xE472e7c3Adc6C752AF49784A229CCf246D3a378a"): true,
		common.HexToAddress("0x4E15F91c562631CAbD856673755700ab19cD8731"): true,
		common.HexToAddress("0x36b3aCe526d18fAccb83888693f0B83263f61Cd1"): true,
		common.HexToAddress("0xEE22F52B73DD00F3D982397Ce4F89fa32D32Ff10"): true,
		common.HexToAddress("0x60756FAd660490F45244d47b75c379F8091E0a76"): true,
		common.HexToAddress("0x38D820E55BEDfbA19c5578Ba6cAe91FD602f9199"): true,
		common.HexToAddress("0xEFF3865b7b414e1b669776289F0a3205bC318A31"): true,
		common.HexToAddress("0x45716dFAA8F92096b6415e389a6b29f896e83c9e"): true,
		common.HexToAddress("0x93Ca620597E4313dEC65B4b6dC21Be071606b043"): true,
		common.HexToAddress("0xBF01e757cEf9f75CBC9ba13282478d1Bf80FAf79"): true,
		common.HexToAddress("0x69C8789288E81F17Bba51D6f34cDbDFd5Ec8eFF5"): true,
		common.HexToAddress("0x5cA8c98F9CE55E760208717EcD51c6ccA57D4516"): true,
		common.HexToAddress("0xB0e554464bb3775198FC45beAa239d5491fB79f4"): true,
		common.HexToAddress("0x7188DC93015ce8AE5E81cfc7361d01cA3B08bD01"): true,
		common.HexToAddress("0xE735dA8BCBC9906324ec7A7c598350ddDbd6F576"): true,
		common.HexToAddress("0xbf8bC28D652e8E35df79815bC6fc1b68C34A445c"): true,
		common.HexToAddress("0x3d6Ee01E5824cf55dD078956FEe6863A43DE4043"): true,
		common.HexToAddress("0xCe213B932E36E6b6E9eF309ffAd0B2982432C65A"): true,
		common.HexToAddress("0x42078855485ADD6d8Edfd5cca710d681fF59Fb6a"): true,
		common.HexToAddress("0xDB3897c45412a027A4A999C96489e630C5dE3060"): true,
		common.HexToAddress("0xE11f55feFc8668510d90A3d7455c965370e1c285"): true,
		common.HexToAddress("0x710d6accE2f2F39BA87e2f9C7c6A8ed73062927e"): true,
		common.HexToAddress("0x5C553bF70FE358b890069a5C240b2EAD3aBDD8D6"): true,
		common.HexToAddress("0xb817ff925Dfc500A81C5b8220dfb1d334782Eef3"): true,
		common.HexToAddress("0x88a56d68d6929F88611944c3FBf25Ce3dbc04292"): true,
		common.HexToAddress("0x39A35E169bE21103C3D2ff224669a0ec928FC115"): true,
		common.HexToAddress("0x2555B09A7Fd10CFD9200C169e039A742680D7dd0"): true,
		common.HexToAddress("0x2D8F648581F1cA008c503e3C2416D9022242e031"): true,
		common.HexToAddress("0x349557A84DE2199a2C4C36423EeC098fbeC515D4"): true,
		common.HexToAddress("0x36934C8675a1eBA3f8c3bCdAE50A4c83ebC46552"): true,
		common.HexToAddress("0x747eC0D69881D63cce7C8007078c69008397DC6c"): true,
		common.HexToAddress("0x30553a22d2aCBc1466fda24208b811e8Ec93cFE9"): true,
		common.HexToAddress("0x8D8b2DDfeE4CD7EF8a6d5A3853118c80C67F7D0E"): true,
		common.HexToAddress("0x7F9CeBdDF507396428dA896D0b089f8bBE37FA8C"): true,
		common.HexToAddress("0x5234fb5748627Fa62179DC84E1721a2775Feb169"): true,
		common.HexToAddress("0x88F68613A298f1f57475F5cfB35721D867Bf5edC"): true,
		common.HexToAddress("0x18a8DbA341149395c571dE6F0eb4559503F24CfD"): true,
		common.HexToAddress("0xCb5136aEE94E977422f13efccee95C3a0b3b32ee"): true,
		common.HexToAddress("0x2ae2Cf6dc19834B77120067553bA203d81218133"): true,
		common.HexToAddress("0xd49987C65e51D62eB5D7759d473823FeD6E1B55e"): true,
		common.HexToAddress("0x71b52a98292CcEE40cdd465363b040cd0b59674c"): true,
		common.HexToAddress("0x5aa1700D6a253CD04aEB048D8114DB593E24986B"): true,
		common.HexToAddress("0x6f7a5b5Be1e4A6A3090F7C42694B7C662814399d"): true,
		common.HexToAddress("0x75297852b7845774db585DDe37219D140C357F42"): true,
		common.HexToAddress("0x95a7627E350D3a17672Be997574Bc986a5944299"): true,
		common.HexToAddress("0x3610a577cBc6ef303B8bE92c1cbc14ee4339c43D"): true,
		common.HexToAddress("0x7451Db9DB860a2857B7C10187eF60633977182ee"): true,
		common.HexToAddress("0x6191FE563815067e4d1De3d382De0557b2106062"): true,
		common.HexToAddress("0x90404e597Ed9Af069452f81ACA764006DaE07bf6"): true,
		common.HexToAddress("0x38ca1E8cB2234Fbca1Fd519Da48d1413E3b1BF32"): true,
		common.HexToAddress("0x24eb91c03eDEe00418Ce3B1fdCA04b36Ff370AfA"): true,
		common.HexToAddress("0x44f59552428efefd84C2a7D49fa95168AE2Dee44"): true,
		common.HexToAddress("0x7Daea33Df029aeD173D4e0a5684F1d67d4892Bb6"): true,
		common.HexToAddress("0x602dd073653584E401cdeda18830e76B0C0F7302"): true,
		common.HexToAddress("0xc8d27709E021d10fDEF8C67b13d8EF133b2118aF"): true,
		common.HexToAddress("0xaC8e7cfEa0fB315916a0a8d506d7942c2faaB0fe"): true,
		common.HexToAddress("0xBc3C35322119F5C46a75011dA9c5BC8820F7DebE"): true,
		common.HexToAddress("0x5fb2379de343EeD7975A8CFD8745441f08648A7d"): true,
		common.HexToAddress("0xFF310da99C8470203425AFE276361D8718cAE021"): true,
		common.HexToAddress("0x462C19154fF6F4C72D0f90982fA714CbF109D30e"): true,
		common.HexToAddress("0x75FF1A46f97eB14162bC544B3c3E7DD67f6693e8"): true,
		common.HexToAddress("0x0e2C16A43518028138EFC0725699973c1F9b7eCe"): true,
		common.HexToAddress("0x2db8636FE90B6D439Ab4Fa86Cb1f8d5f857d966A"): true,
		common.HexToAddress("0xa6A0c7360Ea57cb1dFD4D28904C83899D19b9e60"): true,
		common.HexToAddress("0xeFae9a0f52ea0815CC40E2568ecE94d54d9957Ee"): true,
		common.HexToAddress("0xA96b0cF0311EDfA7CDb3959d9cc133C86555f019"): true,
		common.HexToAddress("0x7dafaED2D1a165d480dc9a5427496C483CBe8cE4"): true,
		common.HexToAddress("0xbCC1090a9c7475F097fFd819ED7119f67957afA3"): true,
		common.HexToAddress("0xAe315BE5Ae8bD11Ec6e650FE4cC64506858781dC"): true,
		common.HexToAddress("0x474c67f458be24977C140c014A54eDc85e13F512"): true,
		common.HexToAddress("0xce365D7CF8845521714A5302bDD3b8f99EA81Bd0"): true,
		common.HexToAddress("0x9E5CD90AA2d4940BD0aC169DecE3fb8638559Dc8"): true,
		common.HexToAddress("0x1f48555529C2633F7e1847d6a48647fb9cCC2b92"): true,
		common.HexToAddress("0x6E7F91E6dbb873ba558f1045604F85eA27933C2B"): true,
		common.HexToAddress("0x59c7795D402fCD39A16E1A3174ADf80E7B229A99"): true,
		common.HexToAddress("0x80Ae8e190ca7471Dae6146558F09577fF120759e"): true,
		common.HexToAddress("0xa85d6e9f49b4383BFce513167Dc18c5911E687f7"): true,
		common.HexToAddress("0x45827B02E64f3106ef219c99397cdBD507543661"): true,
		common.HexToAddress("0x2b49d8dd2974712c932CD3eAE8092c8Dc018aFA9"): true,
		common.HexToAddress("0x8ca070d8c419e4BD053232B49031606EA569E207"): true,
		common.HexToAddress("0x39821b66DcFA940845874706b1013087d9523a44"): true,
		common.HexToAddress("0x0B0E129F5ef0dB88dc6C8d7B1047f8818bd0F73C"): true,
		common.HexToAddress("0xE342D6375ED4821582125D8Ffe1A8925C8B5c8f6"): true,
		common.HexToAddress("0x3b7Ee1CD8c4d56fEf669E927645B9F513652e143"): true,
		common.HexToAddress("0x03e09a71150500Af52A7dfBBC8a18b8902362e36"): true,
		common.HexToAddress("0x31886661046106f0837a602AF8181169BA690a59"): true,
		common.HexToAddress("0x8589dc86EB656afc6b46ca6761f3bBDc7e01FFc5"): true,
		common.HexToAddress("0x8Dbf46Becd5A480d4a6EdE65F5ef279C5afA0F5C"): true,
		common.HexToAddress("0xE800F70e34b3586D0A006ab603fAbC83f68037Cc"): true,
		common.HexToAddress("0x0e7a36Fa55D5CAedA87EEc9E31e925db6D2Ae6c6"): true,
		common.HexToAddress("0x6AEE7FB57e56f1bFcE8A170aaA60b8C41342818f"): true,
		common.HexToAddress("0x5e590cdf9ad38A85d62E036d3f5A78Ad25d9bdE3"): true,
		common.HexToAddress("0x1a8245E97083185ab27971b843DF1B2365360445"): true,
		common.HexToAddress("0xe43F7AB3Dc4719173B7504F88b05eAbDe453312a"): true,
		common.HexToAddress("0x32F1c8818eb2EC082A000630911BFc3c7d0F7dE0"): true,
		common.HexToAddress("0x0505B5c1dE076617dc8dcF36f852577B6905c9F2"): true,
		common.HexToAddress("0x711f0F98426194D1E324da5BDf3F86A3a702A19d"): true,
		common.HexToAddress("0x9fb635B3EcAE30dEC8A17Dc14fd7d694Bf306159"): true,
		common.HexToAddress("0x503672562dF8C1186E3589e21ac20ebDd66699B4"): true,
		common.HexToAddress("0x315C702564cC65f2D4Bc42C66Fd1Be2aafe26f3E"): true,
		common.HexToAddress("0x9058adE67e1f93Abe4D220b77DFbfcBb8f9b0F55"): true,
		common.HexToAddress("0xec3a32c5c5350662BB080923f257D79fb06D4521"): true,
		common.HexToAddress("0xaEaC4E234ba5ED42dE918F11b32E027Bb6EB08Ab"): true,
		common.HexToAddress("0xD22b1CE735264C2cbB09B23BE4fB16D3737c06Be"): true,
		common.HexToAddress("0x67854aD1F8658e82390D4EE952ffA0AC2F2DaA03"): true,
		common.HexToAddress("0x8070F339d7F81545dFC31AB7a8B9720b02C2f6Dc"): true,
		common.HexToAddress("0xEF7Ab7904e1E022A3e6a66A8B66c5fc59a253035"): true,
		common.HexToAddress("0x723a7603129Eba9f212Ee36df6f189323E161C75"): true,
		common.HexToAddress("0xCbe080C6b952aE52eA39c5Fd167af69E3F6F029e"): true,
		common.HexToAddress("0x6e36B057CE652B08867a5494CaE37B01B18F7797"): true,
		common.HexToAddress("0xea9703574790F3Bce4e35E425baa7a0896B1e4eA"): true,
		common.HexToAddress("0x7A27a87FdE7040581B6048b59e26FcbF01769e78"): true,
		common.HexToAddress("0xf561EFB0282bD077AEe4c4Dc02222083D2a3cC79"): true,
		common.HexToAddress("0x4EfC46a907813BCedC7F7b599b6EA3536F1C1CDe"): true,
		common.HexToAddress("0x17e659da1892d26422617983547872beDd64040A"): true,
		common.HexToAddress("0xa50D46b7189537D089388f80eF558892633D95eE"): true,
		common.HexToAddress("0x72127541f77D8B8560fe44425AF35D25ab3ff933"): true,
		common.HexToAddress("0x99e67BA466912E80FFa137E5F5F74C1E65fd8883"): true,
		common.HexToAddress("0x4372876c44D53be8048F736a8596E28fA38162b5"): true,
		common.HexToAddress("0xE4CCf8cF39bb74aB18fCD20d4521b3127cE0D717"): true,
		common.HexToAddress("0x10cBD4D364A47446C38bb7Ffdc9ec6726412420a"): true,
		common.HexToAddress("0xf66E907B60B3cf065d0B4d68Be1d947Ce3dC66a4"): true,
		common.HexToAddress("0x34052cDDA626704231481dD639Be91B97733E8C5"): true,
		common.HexToAddress("0x280cD1E203eC16ec01a7e790934fb9026d3c29F8"): true,
		common.HexToAddress("0x852F8e46449Cd8F84b3Aa87Be248806f9e8168C5"): true,
		common.HexToAddress("0xC8cc2945219850047b5Bacdf73C11a2B492aBc1F"): true,
		common.HexToAddress("0xcCE18C201ADC8E1AE8a2Af1487102D8F6bFd5AE5"): true,
		common.HexToAddress("0x7EA759Ca64a676Cea39E25f32e32f710DE309162"): true,
		common.HexToAddress("0x060e13A85cf30b0E413a16812Dc52935E3F92E74"): true,
		common.HexToAddress("0x1CE20Aec7f19db95D9a7f74FaE4eceD26060Fc3C"): true,
		common.HexToAddress("0xB33673c471F349A628f00D61A44Edb8ef1f97317"): true,
		common.HexToAddress("0xeD2cE76df63bDB6A476fba918507b5EE4835A2Fe"): true,
		common.HexToAddress("0x67688FD14f7fe3F7d748A6c99c7FE57F1501b860"): true,
		common.HexToAddress("0xe442A526049e5498A9d3A18ac79E016D72d384eE"): true,
		common.HexToAddress("0x5BfC7152aedd1E1055381A280C8be1b7aC558456"): true,
		common.HexToAddress("0xA21077C6e8457666cE07a296977192C434544c75"): true,
		common.HexToAddress("0xC86614511be82794BcC47A26914139654884ca7A"): true,
		common.HexToAddress("0x5A53530e96Ca65e7b1477b182Ab29eAd651c78AC"): true,
		common.HexToAddress("0x9E9223A3264D27F00bC62EfBaF512DCe89504a76"): true,
		common.HexToAddress("0x63A36aa6d209E1B604FDA11991A721b479d6e6B0"): true,
		common.HexToAddress("0x7329302ff3AC8B6fC47245Dd4f7Ae1173F9984d5"): true,
		common.HexToAddress("0x1AbA8563b3397E86aa386293921A3BA0f036fC77"): true,
		common.HexToAddress("0xd0EA4dB561b5C1FbddAB282032BC0B2b4D38c4f6"): true,
		common.HexToAddress("0x0588f6032f04033785bBe94B778A7974aCDe8590"): true,
		common.HexToAddress("0x15Bf0D6d063Ee162ac0531Fc7Ba16be8ace6B9ae"): true,
		common.HexToAddress("0xA51D6fB5d883920Cf068FE2dDD61029Ac4596c05"): true,
		common.HexToAddress("0xAA8d0434AAF4b17DDc303E1b0ca350dd1D719d8D"): true,
		common.HexToAddress("0x6dA648B2Ca1B1d98B206c340B085587e834a08f2"): true,
		common.HexToAddress("0xd6dB869D2c2Fc4e1c76daF54D39fA1d89371fcAF"): true,
		common.HexToAddress("0x7c58460dEa812Df4c290F448B4feF0B192cCaAcC"): true,
		common.HexToAddress("0x3BC80b50ff196B055E7f3e13c5DD4b0e1b1766Fc"): true,
		common.HexToAddress("0xe5416ced8Cdc76043d3E98040688720063867A52"): true,
		common.HexToAddress("0x77f89d50eE01351c7E0232303378CCA237e1097c"): true,
		common.HexToAddress("0xCc2c0564E7C2fE0DE19203775e3d5aa81C642782"): true,
		common.HexToAddress("0xe463b08E98Bb9F63358feCC5508939B3c1B972b7"): true,
		common.HexToAddress("0x5C567489AAc47a04d90fa324a10a4A9A995da8E9"): true,
		common.HexToAddress("0xeD286E4135564a5189280d174F5d9262E2143FE7"): true,
		common.HexToAddress("0x49d29CA7423e957Da4Bf74490Cd4ACC36ABC74BD"): true,
		common.HexToAddress("0xa2769F6229E2b713b0807401ce1aA4BaC751BB47"): true,
		common.HexToAddress("0x1a2cF503dc7251E98a5e901ee7E7De8444323Cda"): true,
		common.HexToAddress("0x0c91bC7c7e408d66E66F8Cd77fC1C721a1B100d7"): true,
		common.HexToAddress("0xbDFe226AB6E50ed5547D7fc30A69bE15f2A70017"): true,
		common.HexToAddress("0x551F58526F555c83B66d30F56FDEB9d05A488218"): true,
		common.HexToAddress("0x4c1A1E33C48b3CDa327c87cBd629071B78e2CcAb"): true,
		common.HexToAddress("0xa804Cc7E13326FF402d507e99824EbB417ee66fF"): true,
		common.HexToAddress("0xd8b9Baa99aEA234D893460CdE4B4637968ec215c"): true,
		common.HexToAddress("0x3cd244beFFF7480ebE770ff9755b0E0eedCFe6A1"): true,
		common.HexToAddress("0x8Dc1a19Df25348983B00Ed523a010705ab3E8C07"): true,
		common.HexToAddress("0x90c260F0abDCF8CC324483AC8B02c78D78ceAe32"): true,
		common.HexToAddress("0x8a6A3B47485c007513954EfDa61A449f34716834"): true,
		common.HexToAddress("0xf1abC7aFa493A081558bB8A23e4b945cE1F9038d"): true,
		common.HexToAddress("0x012A74B8C299FB6DC2cC8C19C0d4Dc12dB0447a5"): true,
		common.HexToAddress("0x23e1ec044D3c7d7421379bF31C3f69070d7d268A"): true,
		common.HexToAddress("0xB1F98c4540dCC562158CE814f0D504f6AA73EC19"): true,
		common.HexToAddress("0x89Ef820ED7a0c8f218FDC415c42A29cf4095396A"): true,
		common.HexToAddress("0x207C3d7e401a7482b0e0F6c639BAA193f4981a5f"): true,
		common.HexToAddress("0x8CB35EC9A9CBe9CE1e21c294C5957C16490a911F"): true,
		common.HexToAddress("0xB2715ed78fc9dAe6673ec2866e1f345f91Bc4c7A"): true,
		common.HexToAddress("0xf2430df75B3518fdF432e09E4924E4F74533286E"): true,
		common.HexToAddress("0x4AFa8b6066eCAf6Fc4f41A672D426a18f6B96202"): true,
		common.HexToAddress("0x7Cb9F4100acb12D16e2eD93436c1986A80ebD70c"): true,
		common.HexToAddress("0x7c8780a41A003B0b94A34b0adB7758492E62D76e"): true,
		common.HexToAddress("0x6CAbf4F5cCfC2f8a562C880224de7D3b01225008"): true,
		common.HexToAddress("0x7d353b43C73F6eC24a62Dd3914584300BB5cBB56"): true,
		common.HexToAddress("0xF565AfB440b364a366e126274740190348EDD59a"): true,
		common.HexToAddress("0x31C419eC5005Ab0b3244fa834a5142cbCA5FB81b"): true,
		common.HexToAddress("0xaedbcE5619d4EfFFDF9f43d49C70693129955E33"): true,
		common.HexToAddress("0x4A06dC6Ed9C6faABb9dA4fa74F914876223EE3d6"): true,
		common.HexToAddress("0x102b7aaf41aEbd617487d1E9dD82A07537D5c7dE"): true,
		common.HexToAddress("0x26Fc04d2776b488Ec8aac8f76f45cf4781f0604F"): true,
		common.HexToAddress("0xA08c95FF9A3d5F9739aE69A06f029e9C5050c27D"): true,
		common.HexToAddress("0x2b94A1F42288E4bC16D16FC37d28174529b84c3A"): true,
		common.HexToAddress("0xe3Fcc62B0632e8090976C88D99aB2E418ea63c86"): true,
		common.HexToAddress("0x023Eb42CAA4875e932354a5C00E39e8a36D0FBe3"): true,
		common.HexToAddress("0x1108bEBb9FF16DF01b3DA2a1a0409ec1D4Fa9Bb8"): true,
		common.HexToAddress("0xA2aC4C0Dae004a69370aA54B8932E590CceF069f"): true,
		common.HexToAddress("0x471eb0883840717BA2E0d6AEd085a600a05170de"): true,
		common.HexToAddress("0x688548560DBdFE7672de9104C49112ed6fd11B4D"): true,
		common.HexToAddress("0xC91FBb66804D2E3119DA18939f5A6D918bCcB0f0"): true,
		common.HexToAddress("0xfC92cD0269CC12299f3305A172c0c237321338e0"): true,
		common.HexToAddress("0xEaD0B8aED8307c699f384D847311c9D3961c0ABa"): true,
		common.HexToAddress("0x935Fff7bD2F22AcA2d23E99D92f90b8C7AAcC591"): true,
		common.HexToAddress("0xe44EEBD857CeB4b26199E70b04e4d07267e92Aa9"): true,
		common.HexToAddress("0x47320DeF7c53D5977b311280c5a79b7612ffa938"): true,
		common.HexToAddress("0x7f82a6b7A1043fc8Dc0A4CFEeeB3205c8b665141"): true,
		common.HexToAddress("0xE5FDcF62Ebc48cAa7De15c49100EF147C3026b2d"): true,
		common.HexToAddress("0x05c0a07Ef7Dad1b667AD9A33143D03a6BB7ca266"): true,
		common.HexToAddress("0x6a30c352eca88e8BB5F54b05559c3d023E799aE0"): true,
		common.HexToAddress("0x58A2C6b355C1731DcEB76EF56eE8CB288Dd3E2A1"): true,
		common.HexToAddress("0xaa5D46b6fbd8E5f4D0C4bb59b0c6f08c473cfD6c"): true,
		common.HexToAddress("0x4A0dD578197822F84E8f97aa0c3a2008E8415C2F"): true,
		common.HexToAddress("0xFEB26750D46B70b44E968d9740B4967dE8BBC914"): true,
		common.HexToAddress("0xbe27968ee4a9b9e3d3DDbC04668E417298A7EC27"): true,
		common.HexToAddress("0x4063d0461E2e2EaC48caF529dbB6460cbba6E371"): true,
		common.HexToAddress("0x776b2B6613887DD8C82583CE30f5e9a4f67e6B81"): true,
		common.HexToAddress("0xd25453f7816e0Bc5DcbcC30B3F01673d4c674042"): true,
		common.HexToAddress("0x2314Ea8EfadFc630659C19E42DB2Eff5b723c47B"): true,
		common.HexToAddress("0x3395683BA067E40A6DBEE738Bce56fa64A3705A3"): true,
		common.HexToAddress("0x1170c2474834b3b17DeE8D4785c7F7E09b63cdEc"): true,
		common.HexToAddress("0x45F60CDf5CBe0Eb63f061346281f44C71e3D422d"): true,
		common.HexToAddress("0x1dC0047D0d5d209F6d1Ec2d25E9F260c5eF026d5"): true,
	}
)

func PrintSystemAddress() {
	systemAddr := common.HexToAddress(ValidatorFactoryContract)

	for i := 0; i < 256; i++ {
		addr := crypto.CreateAddress(systemAddr, uint64(i))
		log.Info("systemAddr", "index", i, "addr", addr)
	}
}

func isHubTransition(to *common.Address, data []byte) bool {
	if to == nil || data == nil || len(data) < 4 {
		return false
	}
	systemAddr := common.HexToAddress(ValidatorHubContract)
	if bytes.Compare(to[:], systemAddr[:]) != 0 {
		return false
	}
	return true
}

func IsSlashTransition(to *common.Address, data []byte) bool {
	if isHubTransition(to, data) == false {
		return false
	}
	if len(data) == 36 && hexutil.Encode(data[:4]) == "0xc96be4cb" { //slash(address)
		return true
	}
	return false
}

func IsSyncHeaderTransition(to *common.Address, data []byte) bool {
	if isHubTransition(to, data) == false {
		return false
	}
	if len(data) == 36 && hexutil.Encode(data[:4]) == "0xffd8136e" { //syncTendermintHeader(uint256)
		return true
	}
	return false
}

func IsSystemTransition(to *common.Address, data []byte) bool {
	return IsSlashTransition(to, data) || IsSyncHeaderTransition(to, data)
}
