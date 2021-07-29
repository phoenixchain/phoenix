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
	ValidatorHubContract = "0x0000000000000000000000000000000000001006"
)

var (
	SystemContractAddress = map[common.Address]bool{
		common.HexToAddress("0xea91Ad16d2B6eC90FE255a49Dfd6e2f47304dE88"): true,
		common.HexToAddress("0xACD8D3C8F27ec64d3E8F7Ae5395d33B08Eb29590"): true,
		common.HexToAddress("0x0df21e63812E47fa47765061c3D74dD992e494C2"): true,
		common.HexToAddress("0xED93cC615269127aDD7dfB2A32Ac54AaE346F818"): true,
		common.HexToAddress("0xaA242672590B032557d47AA69841da351593fF67"): true,
		common.HexToAddress("0x07f3f5b5fb0582E76Ab6960a0A925e0cb8585222"): true,
		common.HexToAddress("0xC7A14F2e95A1872Abe772F0B35964d30b86cbbf4"): true,
		common.HexToAddress("0xAb31F53D38c9C5799133668E559B722940B16cbF"): true,
		common.HexToAddress("0x0587105c35079c6fFE4FD9Dc14E6cA63B15b104C"): true,
		common.HexToAddress("0x813BDb7a5062A60a48232CAd68E56855f7f6A54B"): true,
		common.HexToAddress("0x308f400Cb34F46783de73F9e902acC6dee078c3B"): true,
		common.HexToAddress("0x59778A74BFE65391c72a0161eFb618EAD9E56cA7"): true,
		common.HexToAddress("0x9B1ce3076c6D3F887794095aBbC27bA8F12f374F"): true,
		common.HexToAddress("0x50265dDEd979756BC730192e7902AF71A436E9f2"): true,
		common.HexToAddress("0x3CC0E75C419B9Af917F0CB3BFb353aa4558FA158"): true,
		common.HexToAddress("0xbdAb549a283CC0602570994E9B9c394C77332259"): true,
		common.HexToAddress("0xeEDeAd7747d833eE024d2b1749C0edC118693028"): true,
		common.HexToAddress("0xF7aB02ee396CEa74744674D3D6B3236Bf6774eb2"): true,
		common.HexToAddress("0xa0bB62d4Fb037975730C08e43Eb2a666443aC4c4"): true,
		common.HexToAddress("0x525a7e10C79518B6F79642A839b4e70Dc8a74366"): true,
		common.HexToAddress("0xFdf7f2f937abCA615b5C5f917B9E2F5853A5B126"): true,
		common.HexToAddress("0xfA099852440339732Ea6D9799b886F58e845A034"): true,
		common.HexToAddress("0xBc2DB336bdf57bAd38604dB4cCCED629cbD1472B"): true,
		common.HexToAddress("0x1667B5b1d4cA7Fe1719313bA47566bC5bAb7caFA"): true,
		common.HexToAddress("0x9055e8590a258D569819EbbDa2e88982a43dF821"): true,
		common.HexToAddress("0xBc07bdC273B15569ae048AeA5A8B3667D3b743Ce"): true,
		common.HexToAddress("0xFA2E8c5b7c3dcdF1E4A0c15ccC13790A542e5c21"): true,
		common.HexToAddress("0x4948acDA15595B2912aD490067aB72aaA4822B5B"): true,
		common.HexToAddress("0xBA3f6Aeb64cFFfE1458dEf7e08BDA7B002A9E401"): true,
		common.HexToAddress("0xebe96395Cb26Faa360bC87CeAD82316773058749"): true,
		common.HexToAddress("0xC3de7a9324862BAF75241773eEaa534Dc29d812d"): true,
		common.HexToAddress("0x01764d9773090F577aeCF8cA0F1D19Bde65fc6a5"): true,
		common.HexToAddress("0x7d0A212F44095699b10FDF1e8A5200159995968a"): true,
		common.HexToAddress("0x2a689C9D8C9Ac14639486a07e8BF96B863D5B4fF"): true,
		common.HexToAddress("0xcB86E0bA0aB1a857Bb4599ff55fb0277390B725c"): true,
		common.HexToAddress("0x1872Dd910BB46B7A37C1139077506F5C35ac017b"): true,
		common.HexToAddress("0xA167b42B91a88A03D24D4ACD55e496ebdBc00222"): true,
		common.HexToAddress("0x7c728b9bBdA5EE219E9E9D80bf951Ded43f30292"): true,
		common.HexToAddress("0x59243048d9895FF600164228cb5fB8E1496ecC89"): true,
		common.HexToAddress("0xE4283dDa291EF2D4df94457765e7EEeA98d991f3"): true,
		common.HexToAddress("0x852B0B26D28A64cEDF179A7b8ed49a2F1ebc5d6A"): true,
		common.HexToAddress("0x1f5075FeCd64562BC0EbE61762e2D58137639741"): true,
		common.HexToAddress("0x4753805f7859FcD1F5E567c8567a9e473a78CF4A"): true,
		common.HexToAddress("0x6F84d136106F160f1DbdCF3491e1378a07959492"): true,
		common.HexToAddress("0xD14d1dc0df8c3f5F90372ded48ECDb00a9f94C1C"): true,
		common.HexToAddress("0xe027B0F21427f2E3EA2a928dddc6358B34DaDBB0"): true,
		common.HexToAddress("0xDc10B9B63dF971A35c5b1DBfb8D4234220fF51F4"): true,
		common.HexToAddress("0x9b8437A7313B9b1D61DeC598A32f40c88c2119a0"): true,
		common.HexToAddress("0xd8131Ea0a935458F1e38C2cEfdC5a6c087Be456a"): true,
		common.HexToAddress("0xE0c0f72816814236010dF7F53b49c4d325750fe8"): true,
		common.HexToAddress("0xa44eFff2059cC846a098b765139f4aF28FA1ab2B"): true,
		common.HexToAddress("0x582db64cb525c23772fECD7003c5e267FA260383"): true,
		common.HexToAddress("0x6F2254DaeE3c686B80De54A34Fb79E5aFDAE7e36"): true,
		common.HexToAddress("0x2a47b7D164be8B415E7b391C3d14919B7E57cF8F"): true,
		common.HexToAddress("0xc55f226B41A2803116b2aF8Cb4F88cfD4948a481"): true,
		common.HexToAddress("0xF3b925D00F4984C7DAED025B44667af68922Fa8c"): true,
		common.HexToAddress("0xb2C7392E16Ad1D634df57D5cee4fdA5742A56f19"): true,
		common.HexToAddress("0x92d024aA0c92bfa85d0A0D0ee93994d2D59fE991"): true,
		common.HexToAddress("0x4A081742d060d31c14B4E7Bc05F5fe434c1e1284"): true,
		common.HexToAddress("0x0c4aEEAD02Ee5675D4a8B53EFbED4B01834552D5"): true,
		common.HexToAddress("0x6f8F9026eF00cd5929196B610e305584380FB9DA"): true,
		common.HexToAddress("0x77c6B8287CbCB8907c84b48F6f4D3E256BDB5656"): true,
		common.HexToAddress("0xb150096FD110ecd2208Cbb93F51c9569664CF903"): true,
		common.HexToAddress("0x706810eDA07E313928E36851Fb7847140CCEC341"): true,
		common.HexToAddress("0xe4D94197bB22c9df77545267ea7A3D0BA8421d1e"): true,
		common.HexToAddress("0x0BCfFAD809CD84D23A0dde8B5b097605D598e3d2"): true,
		common.HexToAddress("0xAaE155714040119f42CcBC5552c336233bC514f8"): true,
		common.HexToAddress("0x326eCc5593993E040928Dc2796338e00334EaBd0"): true,
		common.HexToAddress("0x63CB31DA88D174D53dC702bfbb0C164857eDf120"): true,
		common.HexToAddress("0x0C1aA77C31F2784C009A9Cc4383E0aFae3E45850"): true,
		common.HexToAddress("0xC3E15A963CCFc0821E818A9bDE1c7a8793A18169"): true,
		common.HexToAddress("0xBF4aC5634bD756066040Bc14dE7ec08E405B6f2D"): true,
		common.HexToAddress("0xEd1e26481A9E545D8902987d47160682fdF119c5"): true,
		common.HexToAddress("0x7aB5388e040Ab369BC13A2254Df3f1a38DA44600"): true,
		common.HexToAddress("0x4aB4464545792B81D97E3CAd9a5ebD2f1F7d4769"): true,
		common.HexToAddress("0xBFFEAFadd1B6A32575EE6F3771264157d3d426c1"): true,
		common.HexToAddress("0xb9C7593b56D2d9324fbe14cB566f2A9CC194606C"): true,
		common.HexToAddress("0xFbf47eCcFD0885e4C39Cb9B4413215A27D46cB5A"): true,
		common.HexToAddress("0xc60C34acc390fE603C9E4451EC76680E7d5bd722"): true,
		common.HexToAddress("0xF152cd1Fcb1A578f53D1D1dEd8714291C4b37812"): true,
		common.HexToAddress("0x89e0a599Da058C7528e87da9BBFAe421d1163CcC"): true,
		common.HexToAddress("0x6Fc0fadBcaD862cE61c4f03A3D4CFCd0D2440784"): true,
		common.HexToAddress("0xA186Be0499FFD698237Cd9Cb8d11e3961991Ec94"): true,
		common.HexToAddress("0x23Edb661c15B3DFa430D37E78848312b20B4b1C0"): true,
		common.HexToAddress("0x2d947f5FC71a057cA3992a6DB6CBf26594E3b690"): true,
		common.HexToAddress("0x07836b29d0BFf7a163daC9e1f52642c95be10e2a"): true,
		common.HexToAddress("0xd9A1E9192F7B967a4CCA229f61f35670797d6867"): true,
		common.HexToAddress("0xEDC1A96cEd68dE8491b5f850fe976FE96D60A236"): true,
		common.HexToAddress("0x4EB4F6ea8207FD41526143ff148EB281c257b9E4"): true,
		common.HexToAddress("0x0DB8dC1886437Cf283d5eEF86aF701B561953CC5"): true,
		common.HexToAddress("0x9990C11EcC19288a2f6E7CA7fFE4b9910fFCE102"): true,
		common.HexToAddress("0xD7215a00dA12dB3471148f9CCBF2c95550c4E104"): true,
		common.HexToAddress("0xc2d498BD3a408FC6e75B6fE78193Ca6f95Fab3c1"): true,
		common.HexToAddress("0xA38e52fC4d64bBb5b36202D30995Eca93ee146d1"): true,
		common.HexToAddress("0xd41c776a07876C019A6AC11595e893502074D1cF"): true,
		common.HexToAddress("0x6bD4a239E25Fe84b18488b3924aCd0d594B9CDa5"): true,
		common.HexToAddress("0x94CE0002beC9DEa92065c01024596ba37C6CC51E"): true,
		common.HexToAddress("0x7961322198ebD3484632a25c31Ebad0F8e8B1aC2"): true,
		common.HexToAddress("0x5c7F548797A5393b842096FF60B48fDb7af77f83"): true,
		common.HexToAddress("0x5BdAA0751f551C9c027185e3Ac54206BEd64b693"): true,
		common.HexToAddress("0x7255423Ae1b3aE53fCFb043D46BDc5134Cb55634"): true,
		common.HexToAddress("0xC67BeA7803C85D19Eb248080A43Dc76AB1524AB9"): true,
		common.HexToAddress("0xe51dC337553fd543320a2b2745e6177F7866CE68"): true,
		common.HexToAddress("0x407A74a5de3571DDe372a5209D339D58C8D83032"): true,
		common.HexToAddress("0xbe080db0d018F60e31b0973317F083eB8b44f0D1"): true,
		common.HexToAddress("0xbE88323E3A542AcD9FE99E28cD5E157D923Dee9d"): true,
		common.HexToAddress("0x3002F18ad6e778D04a1dB6802c078FaA49B28468"): true,
		common.HexToAddress("0x3fd7D8e67f8fE89d95E1468Ac953869207E22Ba4"): true,
		common.HexToAddress("0x7C03F859Fa966074BD21C1B3372F1C684A96b7B2"): true,
		common.HexToAddress("0xd63E06C8D06538C5E793410799392844a6a8ac01"): true,
		common.HexToAddress("0x0F1807bB33962F6A0C7e9015d8Ce2cA7588383fd"): true,
		common.HexToAddress("0x66e1456d954826BDB8200652212aB1153b6063A3"): true,
		common.HexToAddress("0x2635417E8646A4eb051739113A5d3409b695c537"): true,
		common.HexToAddress("0x44F3AEB280DF3292C38F2fd0B1c8BF31d5778c48"): true,
		common.HexToAddress("0x813f6430E8eb66115aDA1B21776248dF34711173"): true,
		common.HexToAddress("0x7008aE81d68FCf6Bb04f0ed67a73d8987664016b"): true,
		common.HexToAddress("0x5Aaf07e671f185F6580b14190D899190A7bbd258"): true,
		common.HexToAddress("0xc0fcF01f8a6Dc797e2ea93e505E8827FD117DfA2"): true,
		common.HexToAddress("0x3Fd4e808A2E839bAeD57D0db82c4c0A2B0E1fc03"): true,
		common.HexToAddress("0xd73cdb1dBFEbc04811fDAE698bC488163c804311"): true,
		common.HexToAddress("0x45B1BE8855C0Af33A4e59f264E295f1f0741D443"): true,
		common.HexToAddress("0x7B7735986820dd107DE1cdF7bDc822CaAc7F918c"): true,
		common.HexToAddress("0x5FA454CD813954Db1C4A52fdeB91d615B26b7AA8"): true,
		common.HexToAddress("0x1251bea50378Ba43f12d364AECCb70E5a7a33743"): true,
		common.HexToAddress("0x7c3f7c569731a228389eC829e8338393ab480B90"): true,
		common.HexToAddress("0xA38aa3A23EE9D5813F38A6CD471d4483Ed3dF3eA"): true,
		common.HexToAddress("0x148C47447E252aFE54bB555e93501e2df2F3409b"): true,
		common.HexToAddress("0x7E7100Cb0f79eaCEA897b1B6cF27Ecd68ccb11AC"): true,
		common.HexToAddress("0x7d9668386a44A9712a7A0141C39D7F05abe340ab"): true,
		common.HexToAddress("0xB3701756A336Eeb1Bda8f504233baf759547fc06"): true,
		common.HexToAddress("0xeD2adE362601915F8d639854b8b404e312aB26e5"): true,
		common.HexToAddress("0x10953655705b2ea4B62Fe079979f73277461D398"): true,
		common.HexToAddress("0xD293C94fD0f4A944bD5c7e45EAa3C1E243F0fe87"): true,
		common.HexToAddress("0x0E45C940d8cb9e6a60E772b58BABD5B8B04E4a46"): true,
		common.HexToAddress("0x24f8e01e2d649062809A4e09bE4280B75B36698D"): true,
		common.HexToAddress("0xfeE9dF43F8ec4DD3979D008913730930B7A499f8"): true,
		common.HexToAddress("0xD846A7eC0E9Ceb57f481c6065DfBc142bdd7Ac3B"): true,
		common.HexToAddress("0x3244B6C9e52a23BD8E23AfBf2fd6A9D073C5b12B"): true,
		common.HexToAddress("0x0B8AB9468486EA689d5f5af27237e19d19688d4A"): true,
		common.HexToAddress("0x8f1EFF16cb679e2fC41FD918d65FB87E1d4A3e35"): true,
		common.HexToAddress("0xf5ecc55191fB5960f3bf42f80E46e274dBA3e3ba"): true,
		common.HexToAddress("0xA2F66c4B903Eac6FCFBCF563dfFC1d88329496BD"): true,
		common.HexToAddress("0xdc2ec3CFD5469be9B5Af06AeF0Dc52737b102e97"): true,
		common.HexToAddress("0x4eb8Fc2a07D52a9FF1471AdC7080674b7cB1372A"): true,
		common.HexToAddress("0xc5b39Cb316ae3A8a6380110EF93C6ABFceb9ea19"): true,
		common.HexToAddress("0x90BC3D6A279E5895789676c47C433E0DF5Ebd805"): true,
		common.HexToAddress("0x7E7Ea8B9e95ac723ed2A5e71336b02131F12177E"): true,
		common.HexToAddress("0x4D218Df1A757ef5d5E45AeC02C8044D2a9FeB8a3"): true,
		common.HexToAddress("0x85018429C8611eC2288C3b81447648E64b680c9e"): true,
		common.HexToAddress("0xbF4F6e15A2B30556EB0119257B4394ceaa59d0ac"): true,
		common.HexToAddress("0x170E69a80f003f5f9bB7577aaF27ff1f64Afe40e"): true,
		common.HexToAddress("0x481E5ADAE5BB609Ec30b993E28Feb1d90f1F23e5"): true,
		common.HexToAddress("0x8B9a3eECb28E9eADDEdf3c98C70A86961942681e"): true,
		common.HexToAddress("0x4C0D7129913Be5dDf02B0C269676F9454b843CB2"): true,
		common.HexToAddress("0x0C2092ca3d3f0AB22A668080EaAf1F144e6f7e4B"): true,
		common.HexToAddress("0xF4A1d86584909d996e0225c51E1aD741f854775C"): true,
		common.HexToAddress("0x319c507b9Cc69fA0CeC943CE3FC5C44B8a0B6D33"): true,
		common.HexToAddress("0x3839753775Ad37023F59fc3fe12713e1Ea330048"): true,
		common.HexToAddress("0x20F9F5a771D669F8AE48F85091ac0A409030Ba02"): true,
		common.HexToAddress("0x85D524a6698004D652e604777269086Ca498AEbC"): true,
		common.HexToAddress("0xb1b8D4be7A023206A5AC6cff092588Fa889952EB"): true,
		common.HexToAddress("0x8d4baD091D8b1Fb29f15B8F0b80FBd40a8E4850a"): true,
		common.HexToAddress("0x23ED8320180e3a662359b69DbAd98eb4D1875D69"): true,
		common.HexToAddress("0xa7a968a50DBfbC20CCebB6E20d9DA31dA9eCB281"): true,
		common.HexToAddress("0x2C8c81C8Ed13649c0C0babcAfC6E32802424dCcF"): true,
		common.HexToAddress("0xe70FbD60eCC04A6D59FB18636356264eA4E7d8CB"): true,
		common.HexToAddress("0x4c4496742f1EC89Da7bE36c2d1bCd4eFcbf9efBf"): true,
		common.HexToAddress("0xd42816522333B43a4a78408773f173907d13A89b"): true,
		common.HexToAddress("0x12849E69F11A3d5B45959d55B76BE3d17b185f4a"): true,
		common.HexToAddress("0xa2c46e390CeB8e0ECfea17De73AF0E34716c5912"): true,
		common.HexToAddress("0xa71B93c6963c0E0cB309c521C048e1Fa42F5Dc9B"): true,
		common.HexToAddress("0xe4dBcbe873f3389A8D48fA03d21C7127B6f7Aa9f"): true,
		common.HexToAddress("0xa953044311005B6F7763609f5247DB3DADf874f5"): true,
		common.HexToAddress("0x8b93695ee35A7746C47f80C7C803BA42593Ad445"): true,
		common.HexToAddress("0xC82ee8444fB1a9464824093eF8cBb7Ca41908Bf4"): true,
		common.HexToAddress("0xB3A6beD7d7Ae8695b836Acc9fD8BbA02012Ff639"): true,
		common.HexToAddress("0xf91AcaC82b6E47F564F0eCF6Af00983162a5c07a"): true,
		common.HexToAddress("0x271260b110980bb73085BbdE4ead457DA7D993E0"): true,
		common.HexToAddress("0xe5936F44ce263898d543B66A8C8DDa47A263014d"): true,
		common.HexToAddress("0xD2Dc67294422D4f45A740a8CB82444D1DC7e2dBE"): true,
		common.HexToAddress("0xF1649E811483244DAA903cfCa5c7157b6b19ec7c"): true,
		common.HexToAddress("0x5FbFC86Ae69Ba5f32e891b1Bd55942f620ec37d4"): true,
		common.HexToAddress("0x479A11D79fd58b2E4Bb9A7159702De92728C136D"): true,
		common.HexToAddress("0x124bb22F5A10D53247A5a27B702869fE7EdF7477"): true,
		common.HexToAddress("0x8e75B97c89Ca5ecF3D470E588d08f7953A0f244d"): true,
		common.HexToAddress("0xDdEf89e043F40031396d2f6252d6BC922f0a72F1"): true,
		common.HexToAddress("0xc14CCf2D56F71FC3D88ef111C4cd3af3FDCce6Aa"): true,
		common.HexToAddress("0xEFeed088E9Ba8490504DA4881aa502e961b634B5"): true,
		common.HexToAddress("0x84E6c70823649A6DCafF0aAF1f27103258C952f6"): true,
		common.HexToAddress("0x5A26dfe6C8fa28f869CCf6406f65b2c0D94f2380"): true,
		common.HexToAddress("0xDf5D93852cF1da556625eA914af9A11f6a5D85Ec"): true,
		common.HexToAddress("0xd4B6343219D44383A63e64C574bB3ae9deE2d0c4"): true,
		common.HexToAddress("0xaDBb8ED55C75D0aE34403b527B545E8E21531E95"): true,
		common.HexToAddress("0xb9966D609c0834E2BD534DA416c019861eCfb148"): true,
		common.HexToAddress("0xD19536C07ff947feF4f9400917e22cAd79E986DD"): true,
		common.HexToAddress("0x3Ef79438CF93a9992816DB2Fc647C3f3D2429B90"): true,
		common.HexToAddress("0xA81eB2d83F7b0b1E6aEcb8965A94C975f4bbc930"): true,
		common.HexToAddress("0xc02d613626863A42FaFAebd85869F17127e7d292"): true,
		common.HexToAddress("0x62bbB209f6bc571e56AAfc49fca774672ecDB55c"): true,
		common.HexToAddress("0x277611B0d5073226Ce1B6e8C2D480610f7a393EF"): true,
		common.HexToAddress("0x4467255FA773355690E434fF9e1259A22E531538"): true,
		common.HexToAddress("0x2Cf955C21C4ec23F8dd8eA9388461B4b4B0E9BCe"): true,
		common.HexToAddress("0xF8009dc60AeAfD71D7fBFA05F1c5172181a69Ca1"): true,
		common.HexToAddress("0xE1460b1B3F71150D661f0C8e509287a7c7eD2d5b"): true,
		common.HexToAddress("0x3B55A66f53Cd26123D7f1186e6935629244ef06D"): true,
		common.HexToAddress("0x75f2d035D8e05799F22389d99053BD3b0A464D9c"): true,
		common.HexToAddress("0x4e5EeF7E66ddAdDe095bA69ed878B0F02ed45739"): true,
		common.HexToAddress("0x6888D305F2AA39c9DDFdC2E0c08ff326f5e707b4"): true,
		common.HexToAddress("0x15e7F62094a5bf9cD0197299B305A1e6A4a73527"): true,
		common.HexToAddress("0x076574F373b536E17B88b4691899B0CEbdFBF5C0"): true,
		common.HexToAddress("0xc79ABD943792Fb7E89C381De587278fC006811D8"): true,
		common.HexToAddress("0xa1E6ca22FD5Bc9F78c83580DA44ad2E8d9f907C9"): true,
		common.HexToAddress("0x9ca44ED1d607917D7d5e85fF7A44BBA417E9f5e0"): true,
		common.HexToAddress("0x70939e43b56D7267da9663d5c412A23678a7afB2"): true,
		common.HexToAddress("0x8818ff60A8EC2B4b2460507784aDa8b8905b664a"): true,
		common.HexToAddress("0x9437C8a8F142cBDC7aBB37b66c53Ba7AD74A6EE9"): true,
		common.HexToAddress("0x5292dA617D7c744c9D5930e7D245D75ed8A00875"): true,
		common.HexToAddress("0x7F515301Caa24341ffCA5122f8BFebE4a47AC806"): true,
		common.HexToAddress("0xb773B606E918d2DF55A11e9c58F42eDf2cF0C3C6"): true,
		common.HexToAddress("0x1f3391230412EA07a8C0d35306E4b6f5706a5755"): true,
		common.HexToAddress("0xc03fe034c1453839e0dbA59D1A8f793EdA692acd"): true,
		common.HexToAddress("0xD8359819a88aB10c204c761265C39F732249e2C3"): true,
		common.HexToAddress("0xF34028471D238c5fd8954564E91F02Ae21946F4D"): true,
		common.HexToAddress("0x8F6D4C189087741F18A02538a80925881c897dfF"): true,
		common.HexToAddress("0xd8Cc7E2d1f587c11477269CC03f5c629E0aF96Dd"): true,
		common.HexToAddress("0x3524D9AE14Ddd2cb2f5D97e2f227a40649d672DD"): true,
		common.HexToAddress("0x9b18D8db7aF42C477877b35118dA5818c29678aB"): true,
		common.HexToAddress("0x8c114B196e631Ea6aa940d28Bca00B87A679072B"): true,
		common.HexToAddress("0xA0d802f3b766F1AeB498aFcD0470cDC8fae06DbF"): true,
		common.HexToAddress("0xE6557eD2b294c406871a1e11641E82cf99edD8F7"): true,
		common.HexToAddress("0x5930DFe1a2A310A9698aac65DE68608C914d4e4e"): true,
		common.HexToAddress("0x1b2Bb6552c52b289e381f097eE5A63FE4f3069F8"): true,
		common.HexToAddress("0x388E11321C472cE095a2d4fE13302dDf62AEBbD7"): true,
		common.HexToAddress("0x1c4A5Ca0EEe1e257f66D7ed74C0bF11878Ee2b5D"): true,
		common.HexToAddress("0xE9c952A9ad77FCD007cc82aED2A6Cd3dBf7baD76"): true,
		common.HexToAddress("0xeF00F1cEd5B76b0167C3aB4684D9ED2A01B0F341"): true,
		common.HexToAddress("0x620ae274D8edD6476026F7d95C47331d901f407e"): true,
		common.HexToAddress("0xd61D46503f59FF0ecf84F6b4C9972d364c8137A8"): true,
		common.HexToAddress("0xb405d54A63A19f3A0A0B742715543432BF227516"): true,
		common.HexToAddress("0x1Df6853d44a90546595A32809d55Ba5327Feec53"): true,
		common.HexToAddress("0x8acA295f13C4F339b3B9D05BCBd87Fd1E3Add952"): true,
		common.HexToAddress("0x5f4AB6d21B0823FdaFFfe60bf91CB2C337855d76"): true,
		common.HexToAddress("0xD41E8679fa8BBC6A5d8805Fb6e619C7e3B8E4c85"): true,
		common.HexToAddress("0x3abbA1E71BcBDdcbCF5CFf725812327eAb7F2186"): true,
		common.HexToAddress("0xf1D404eec262204753dEd813f3a5f9Fbb515779E"): true,
		common.HexToAddress("0x6d1B768Dbc886cb3eF57B1E2009890065e32B5ea"): true,
		common.HexToAddress("0x71E56182CB762C67486c6E5DcD761aCFe93fBeBE"): true,
		common.HexToAddress("0x57C2d747B7Bd5E60E909D0F2185200b2F1a15056"): true,
		common.HexToAddress("0x9b28A3BbCB3a877CFF7987800354cD394EDA5211"): true,
		common.HexToAddress("0x7F2E73BAde3F5d26c8367B0F845d03165543076e"): true,
		common.HexToAddress("0x555D98Ef669d476c249E918eC4E9De3c16988027"): true,
		common.HexToAddress("0x1cf489847693093aFAB08B23d2a99615223ccd47"): true,
		common.HexToAddress("0x1AF6F47FE6793663e0BF8A53Bd303404a7986047"): true,
		common.HexToAddress("0xA07936CE3D427820B64d395Af8DB798EF0d960AA"): true,
		common.HexToAddress("0x2aC9f532780d1757d12e8239711992Eea2563056"): true,
		common.HexToAddress("0xcE22d729db4Ea0dc51D721beE1e6d048E778FF64"): true,
	}
)

func PrintSystemAddress() {
	systemAddr := common.HexToAddress(ValidatorHubContract)

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
