package ast_test

import (
	"fmt"
	"os"
	t "testing"

	"github.com/klvnptr/k/testing"

	"github.com/stretchr/testify/suite"
)

type KTestSuite struct {
	testing.CompilerSuite
}

func (suite *KTestSuite) SetupTest() {
	suite.T().Parallel()
}

func (suite *KTestSuite) EqualTestCase(num int) {
	srcPath := fmt.Sprintf("../testsuite/%05d.k", num)
	expectedPath := fmt.Sprintf("../testsuite/%05d.k.expected", num)

	// read source file
	bsrc, err := os.ReadFile(srcPath)
	if err != nil {
		suite.NoError(err)
		return
	}

	// read expected output
	bexpected, err := os.ReadFile(expectedPath)
	if err != nil {
		suite.NoError(err)
		return
	}

	suite.EqualProgramK(string(bsrc), string(bexpected))
}

// file://./../testsuite/00001.k
func (suite *KTestSuite) TestK00001() {
	suite.EqualTestCase(1)
}

// file://./../testsuite/00002.k
func (suite *KTestSuite) TestK00002() {
	suite.EqualTestCase(2)
}

// file://./../testsuite/00003.k
func (suite *KTestSuite) TestK00003() {
	suite.EqualTestCase(3)
}

// file://./../testsuite/00004.k
func (suite *KTestSuite) TestK00004() {
	suite.EqualTestCase(4)
}

// file://./../testsuite/00005.k
func (suite *KTestSuite) TestK00005() {
	suite.EqualTestCase(5)
}

// file://./../testsuite/00006.k
func (suite *KTestSuite) TestK00006() {
	suite.EqualTestCase(6)
}

// file://./../testsuite/00007.k
// func (suite *KTestSuite) TestK00007() {
// 	suite.EqualTestCase(7)
// }

// file://./../testsuite/00008.k
// func (suite *KTestSuite) TestK00008() {
// 	suite.EqualTestCase(8)
// }

// file://./../testsuite/00009.k
func (suite *KTestSuite) TestK00009() {
	suite.EqualTestCase(9)
}

// file://./../testsuite/00010.k
// func (suite *KTestSuite) TestK00010() {
// 	suite.EqualTestCase(10)
// }

// file://./../testsuite/00011.k
// func (suite *KTestSuite) TestK00011() {
// 	suite.EqualTestCase(11)
// }

// file://./../testsuite/00012.k
func (suite *KTestSuite) TestK00012() {
	suite.EqualTestCase(12)
}

// file://./../testsuite/00013.k
func (suite *KTestSuite) TestK00013() {
	suite.EqualTestCase(13)
}

// file://./../testsuite/00014.k
func (suite *KTestSuite) TestK00014() {
	suite.EqualTestCase(14)
}

// file://./../testsuite/00015.k
func (suite *KTestSuite) TestK00015() {
	suite.EqualTestCase(15)
}

// file://./../testsuite/00016.k
func (suite *KTestSuite) TestK00016() {
	suite.EqualTestCase(16)
}

// file://./../testsuite/00017.k
func (suite *KTestSuite) TestK00017() {
	suite.EqualTestCase(17)
}

// file://./../testsuite/00018.k
func (suite *KTestSuite) TestK00018() {
	suite.EqualTestCase(18)
}

// file://./../testsuite/00019.k
func (suite *KTestSuite) TestK00019() {
	suite.EqualTestCase(19)
}

// file://./../testsuite/00020.k
func (suite *KTestSuite) TestK00020() {
	suite.EqualTestCase(20)
}

// file://./../testsuite/00021.k
func (suite *KTestSuite) TestK00021() {
	suite.EqualTestCase(21)
}

// file://./../testsuite/00022.k
func (suite *KTestSuite) TestK00022() {
	suite.EqualTestCase(22)
}

// globals
// file://./../testsuite/00023.k
// func (suite *KTestSuite) TestK00023() {
// 	suite.EqualTestCase(23)
// }

// globals
// file://./../testsuite/00024.k
// func (suite *KTestSuite) TestK00024() {
// 	suite.EqualTestCase(24)
// }

// file://./../testsuite/00025.k
func (suite *KTestSuite) TestK00025() {
	suite.EqualTestCase(25)
}

// file://./../testsuite/00026.k
func (suite *KTestSuite) TestK00026() {
	suite.EqualTestCase(26)
}

// file://./../testsuite/00027.k
func (suite *KTestSuite) TestK00027() {
	suite.EqualTestCase(27)
}

// file://./../testsuite/00028.k
func (suite *KTestSuite) TestK00028() {
	suite.EqualTestCase(28)
}

// file://./../testsuite/00029.k
func (suite *KTestSuite) TestK00029() {
	suite.EqualTestCase(29)
}

// file://./../testsuite/00030.k
func (suite *KTestSuite) TestK00030() {
	suite.EqualTestCase(30)
}

// file://./../testsuite/00031.k
func (suite *KTestSuite) TestK00031() {
	suite.EqualTestCase(31)
}

// file://./../testsuite/00032.k
func (suite *KTestSuite) TestK00032() {
	suite.EqualTestCase(32)
}

// globals
// file://./../testsuite/00033.k
// func (suite *KTestSuite) TestK00033() {
// 	suite.EqualTestCase(33)
// }

// for, continue, break
// file://./../testsuite/00034.k
// func (suite *KTestSuite) TestK00034() {
// 	suite.EqualTestCase(34)
// }

// file://./../testsuite/00035.k
func (suite *KTestSuite) TestK00035() {
	suite.EqualTestCase(35)
}

// +=, -=, *=, /= operators
// file://./../testsuite/00036.k
// func (suite *KTestSuite) TestK00036() {
// 	suite.EqualTestCase(36)
// }

// subtract a pointer from another pointer
// file://./../testsuite/00037.k
// func (suite *KTestSuite) TestK00037() {
// 	suite.EqualTestCase(37)
// }

// sizeof operator
// file://./../testsuite/00038.k
func (suite *KTestSuite) TestK00038() {
	suite.EqualTestCase(38)
}

// file://./../testsuite/00039.k
func (suite *KTestSuite) TestK00039() {
	suite.EqualTestCase(39)
}

// globals
// file://./../testsuite/00040.k
// func (suite *KTestSuite) TestK00040() {
// 	suite.EqualTestCase(40)
// }

// file://./../testsuite/00041.k
func (suite *KTestSuite) TestK00041() {
	suite.EqualTestCase(41)
}

// union
// file://./../testsuite/00042.k
// func (suite *KTestSuite) TestK00042() {
// 	suite.EqualTestCase(42)
// }

// file://./../testsuite/00043.k
func (suite *KTestSuite) TestK00043() {
	suite.EqualTestCase(43)
}

// fancy structs, not gonna do this
// file://./../testsuite/00044.k
// func (suite *KTestSuite) TestK00044() {
// 	suite.EqualTestCase(44)
// }

// globals with init
// file://./../testsuite/00045.k
// func (suite *KTestSuite) TestK00045() {
// 	suite.EqualTestCase(45)
// }

// union, unnamd struct & union
// file://./../testsuite/00046.k
// func (suite *KTestSuite) TestK00046() {
// 	suite.EqualTestCase(46)
// }

// struct assign
// file://./../testsuite/00047.k
// func (suite *KTestSuite) TestK00047() {
// 	suite.EqualTestCase(47)
// }

// struct assign
// file://./../testsuite/00048.k
// func (suite *KTestSuite) TestK00048() {
// 	suite.EqualTestCase(48)
// }

// globals, struct assign
// file://./../testsuite/00049.k
// func (suite *KTestSuite) TestK00049() {
// 	suite.EqualTestCase(49)
// }

// struct assign, unnamed union
// file://./../testsuite/00050.k
// func (suite *KTestSuite) TestK00050() {
// 	suite.EqualTestCase(50)
// }

// switch, label
// file://./../testsuite/00051.k
// func (suite *KTestSuite) TestK00051() {
// 	suite.EqualTestCase(51)
// }

// struct scoping
// file://./../testsuite/00052.k
// func (suite *KTestSuite) TestK00052() {
// 	suite.EqualTestCase(52)
// }

// struct scoping
// file://./../testsuite/00053.k
// func (suite *KTestSuite) TestK00053() {
// 	suite.EqualTestCase(53)
// }

// enum
// file://./../testsuite/00054.k
// func (suite *KTestSuite) TestK00054() {
// 	suite.EqualTestCase(54)
// }

// enum
// file://./../testsuite/00055.k
// func (suite *KTestSuite) TestK00055() {
// 	suite.EqualTestCase(55)
// }

// file://./../testsuite/00056.k
func (suite *KTestSuite) TestK00056() {
	suite.EqualTestCase(56)
}

// file://./../testsuite/00057.k
func (suite *KTestSuite) TestK00057() {
	suite.EqualTestCase(57)
}

// split string literal: "abc" "def"
// file://./../testsuite/00058.k
// func (suite *KTestSuite) TestK00058() {
// 	suite.EqualTestCase(58)
// }

// file://./../testsuite/00059.k
func (suite *KTestSuite) TestK00059() {
	suite.EqualTestCase(59)
}

// file://./../testsuite/00060.k
func (suite *KTestSuite) TestK00060() {
	suite.EqualTestCase(60)
}

// macro
// file://./../testsuite/00061.k
// func (suite *KTestSuite) TestK00061() {
// 	suite.EqualTestCase(61)
// }

// macro
// file://./../testsuite/00062.k
// func (suite *KTestSuite) TestK00062() {
// 	suite.EqualTestCase(62)
// }

// macro
// file://./../testsuite/00063.k
// func (suite *KTestSuite) TestK00063() {
// 	suite.EqualTestCase(63)
// }

// macro
// file://./../testsuite/00064.k
// func (suite *KTestSuite) TestK00064() {
// 	suite.EqualTestCase(64)
// }

// macro
// file://./../testsuite/00065.k
// func (suite *KTestSuite) TestK00065() {
// 	suite.EqualTestCase(65)
// }

// macro
// file://./../testsuite/00066.k
// func (suite *KTestSuite) TestK00066() {
// 	suite.EqualTestCase(66)
// }

// macro
// file://./../testsuite/00067.k
// func (suite *KTestSuite) TestK00067() {
// 	suite.EqualTestCase(67)
// }

// macro
// file://./../testsuite/00068.k
// func (suite *KTestSuite) TestK00068() {
// 	suite.EqualTestCase(68)
// }

// macro
// file://./../testsuite/00069.k
// func (suite *KTestSuite) TestK00069() {
// 	suite.EqualTestCase(69)
// }

// macro
// file://./../testsuite/00070.k
// func (suite *KTestSuite) TestK00070() {
// 	suite.EqualTestCase(70)
// }

// macro
// file://./../testsuite/00071.k
// func (suite *KTestSuite) TestK00071() {
// 	suite.EqualTestCase(71)
// }

// +=, -=, *=, /= operators
// file://./../testsuite/00072.k
func (suite *KTestSuite) TestK00072() {
	suite.EqualTestCase(72)
}

// +=, -=, *=, /= operators
// file://./../testsuite/00073.k
func (suite *KTestSuite) TestK00073() {
	suite.EqualTestCase(73)
}

// macro
// file://./../testsuite/00074.k
// func (suite *KTestSuite) TestK00074() {
// 	suite.EqualTestCase(74)
// }

// macro
// file://./../testsuite/00075.k
// func (suite *KTestSuite) TestK00075() {
// 	suite.EqualTestCase(75)
// }

// inline if expression
// file://./../testsuite/00076.k
// func (suite *KTestSuite) TestK00076() {
// 	suite.EqualTestCase(76)
// }

// file://./../testsuite/00077.k
func (suite *KTestSuite) TestK00077() {
	suite.EqualTestCase(77)
}

// function pointer type
// file://./../testsuite/00078.k
// func (suite *KTestSuite) TestK00078() {
// 	suite.EqualTestCase(78)
// }

// macro
// file://./../testsuite/00079.k
// func (suite *KTestSuite) TestK00079() {
// 	suite.EqualTestCase(79)
// }

// file://./../testsuite/00080.k
func (suite *KTestSuite) TestK00080() {
	suite.EqualTestCase(80)
}

// file://./../testsuite/00081.k
func (suite *KTestSuite) TestK00081() {
	suite.EqualTestCase(81)
}

// file://./../testsuite/00082.k
func (suite *KTestSuite) TestK00082() {
	suite.EqualTestCase(82)
}

// macro
// file://./../testsuite/00083.k
// func (suite *KTestSuite) TestK00083() {
// 	suite.EqualTestCase(83)
// }

// macro
// file://./../testsuite/00084.k
// func (suite *KTestSuite) TestK00084() {
// 	suite.EqualTestCase(84)
// }

// macro
// file://./../testsuite/00085.k
// func (suite *KTestSuite) TestK00085() {
// 	suite.EqualTestCase(85)
// }

// file://./../testsuite/00086.k
func (suite *KTestSuite) TestK00086() {
	suite.EqualTestCase(86)
}

// function pointer type
// file://./../testsuite/00087.k
// func (suite *KTestSuite) TestK00087() {
// 	suite.EqualTestCase(87)
// }

// function pointer type
// file://./../testsuite/00088.k
// func (suite *KTestSuite) TestK00088() {
// 	suite.EqualTestCase(88)
// }

// function pointer type
// file://./../testsuite/00089.k
// func (suite *KTestSuite) TestK00089() {
// 	suite.EqualTestCase(89)
// }

// globals, array & struct assign
// file://./../testsuite/00090.k
// func (suite *KTestSuite) TestK00090() {
// 	suite.EqualTestCase(90)
// }

// struct assign
// file://./../testsuite/00091.k
// func (suite *KTestSuite) TestK00091() {
// 	suite.EqualTestCase(91)
// }

// struct assign
// file://./../testsuite/00092.k
// func (suite *KTestSuite) TestK00092() {
// 	suite.EqualTestCase(92)
// }

// struct assign
// file://./../testsuite/00093.k
// func (suite *KTestSuite) TestK00093() {
// 	suite.EqualTestCase(93)
// }

// extern, multi file
// file://./../testsuite/00094.k
// func (suite *KTestSuite) TestK00094() {
// 	suite.EqualTestCase(94)
// }

// extern, function pointer
// file://./../testsuite/00095.k
// func (suite *KTestSuite) TestK00095() {
// 	suite.EqualTestCase(95)
// }

// globals
// file://./../testsuite/00096.k
// func (suite *KTestSuite) TestK00096() {
// 	suite.EqualTestCase(96)
// }

// macro
// file://./../testsuite/00097.k
// func (suite *KTestSuite) TestK00097() {
// 	suite.EqualTestCase(97)
// }

// file://./../testsuite/00098.k
// func (suite *KTestSuite) TestK00098() {
// 	suite.EqualTestCase(98)
// }

// file://./../testsuite/00099.k
// func (suite *KTestSuite) TestK00099() {
// 	suite.EqualTestCase(99)
// }

// file://./../testsuite/00100.k
func (suite *KTestSuite) TestK00100() {
	suite.EqualTestCase(100)
}

// do-while
// file://./../testsuite/00101.k
// func (suite *KTestSuite) TestK00101() {
// 	suite.EqualTestCase(101)
// }

// file://./../testsuite/00102.k
// func (suite *KTestSuite) TestK00102() {
// 	suite.EqualTestCase(102)
// }

// file://./../testsuite/00103.k
// func (suite *KTestSuite) TestK00103() {
// 	suite.EqualTestCase(103)
// }

// file://./../testsuite/00104.k
// func (suite *KTestSuite) TestK00104() {
// 	suite.EqualTestCase(104)
// }

// file://./../testsuite/00105.k
// func (suite *KTestSuite) TestK00105() {
// 	suite.EqualTestCase(105)
// }

// file://./../testsuite/00106.k
// func (suite *KTestSuite) TestK00106() {
// 	suite.EqualTestCase(106)
// }

// file://./../testsuite/00107.k
// func (suite *KTestSuite) TestK00107() {
// 	suite.EqualTestCase(107)
// }

// file://./../testsuite/00108.k
// func (suite *KTestSuite) TestK00108() {
// 	suite.EqualTestCase(108)
// }

// file://./../testsuite/00109.k
// func (suite *KTestSuite) TestK00109() {
// 	suite.EqualTestCase(109)
// }

// file://./../testsuite/00110.k
// func (suite *KTestSuite) TestK00110() {
// 	suite.EqualTestCase(110)
// }

// file://./../testsuite/00111.k
// func (suite *KTestSuite) TestK00111() {
// 	suite.EqualTestCase(111)
// }

// file://./../testsuite/00112.k
// func (suite *KTestSuite) TestK00112() {
// 	suite.EqualTestCase(112)
// }

// file://./../testsuite/00113.k
// func (suite *KTestSuite) TestK00113() {
// 	suite.EqualTestCase(113)
// }

// file://./../testsuite/00114.k
// func (suite *KTestSuite) TestK00114() {
// 	suite.EqualTestCase(114)
// }

// file://./../testsuite/00115.k
// func (suite *KTestSuite) TestK00115() {
// 	suite.EqualTestCase(115)
// }

// file://./../testsuite/00116.k
// func (suite *KTestSuite) TestK00116() {
// 	suite.EqualTestCase(116)
// }

// file://./../testsuite/00117.k
// func (suite *KTestSuite) TestK00117() {
// 	suite.EqualTestCase(117)
// }

// file://./../testsuite/00118.k
// func (suite *KTestSuite) TestK00118() {
// 	suite.EqualTestCase(118)
// }

// file://./../testsuite/00119.k
// func (suite *KTestSuite) TestK00119() {
// 	suite.EqualTestCase(119)
// }

// file://./../testsuite/00120.k
// func (suite *KTestSuite) TestK00120() {
// 	suite.EqualTestCase(120)
// }

// file://./../testsuite/00121.k
// func (suite *KTestSuite) TestK00121() {
// 	suite.EqualTestCase(121)
// }

// file://./../testsuite/00122.k
// func (suite *KTestSuite) TestK00122() {
// 	suite.EqualTestCase(122)
// }

// file://./../testsuite/00123.k
// func (suite *KTestSuite) TestK00123() {
// 	suite.EqualTestCase(123)
// }

// file://./../testsuite/00124.k
// func (suite *KTestSuite) TestK00124() {
// 	suite.EqualTestCase(124)
// }

// file://./../testsuite/00125.k
// func (suite *KTestSuite) TestK00125() {
// 	suite.EqualTestCase(125)
// }

// file://./../testsuite/00126.k
// func (suite *KTestSuite) TestK00126() {
// 	suite.EqualTestCase(126)
// }

// file://./../testsuite/00127.k
// func (suite *KTestSuite) TestK00127() {
// 	suite.EqualTestCase(127)
// }

// file://./../testsuite/00128.k
// func (suite *KTestSuite) TestK00128() {
// 	suite.EqualTestCase(128)
// }

// file://./../testsuite/00129.k
// func (suite *KTestSuite) TestK00129() {
// 	suite.EqualTestCase(129)
// }

// file://./../testsuite/00130.k
// func (suite *KTestSuite) TestK00130() {
// 	suite.EqualTestCase(130)
// }

// file://./../testsuite/00131.k
// func (suite *KTestSuite) TestK00131() {
// 	suite.EqualTestCase(131)
// }

// file://./../testsuite/00132.k
// func (suite *KTestSuite) TestK00132() {
// 	suite.EqualTestCase(132)
// }

// file://./../testsuite/00133.k
// func (suite *KTestSuite) TestK00133() {
// 	suite.EqualTestCase(133)
// }

// file://./../testsuite/00134.k
// func (suite *KTestSuite) TestK00134() {
// 	suite.EqualTestCase(134)
// }

// file://./../testsuite/00135.k
// func (suite *KTestSuite) TestK00135() {
// 	suite.EqualTestCase(135)
// }

// file://./../testsuite/00136.k
// func (suite *KTestSuite) TestK00136() {
// 	suite.EqualTestCase(136)
// }

// file://./../testsuite/00137.k
// func (suite *KTestSuite) TestK00137() {
// 	suite.EqualTestCase(137)
// }

// file://./../testsuite/00138.k
// func (suite *KTestSuite) TestK00138() {
// 	suite.EqualTestCase(138)
// }

// file://./../testsuite/00139.k
// func (suite *KTestSuite) TestK00139() {
// 	suite.EqualTestCase(139)
// }

// file://./../testsuite/00140.k
// func (suite *KTestSuite) TestK00140() {
// 	suite.EqualTestCase(140)
// }

// file://./../testsuite/00141.k
// func (suite *KTestSuite) TestK00141() {
// 	suite.EqualTestCase(141)
// }

// file://./../testsuite/00142.k
// func (suite *KTestSuite) TestK00142() {
// 	suite.EqualTestCase(142)
// }

// file://./../testsuite/00143.k
// func (suite *KTestSuite) TestK00143() {
// 	suite.EqualTestCase(143)
// }

// file://./../testsuite/00144.k
// func (suite *KTestSuite) TestK00144() {
// 	suite.EqualTestCase(144)
// }

// file://./../testsuite/00145.k
// func (suite *KTestSuite) TestK00145() {
// 	suite.EqualTestCase(145)
// }

// file://./../testsuite/00146.k
// func (suite *KTestSuite) TestK00146() {
// 	suite.EqualTestCase(146)
// }

// file://./../testsuite/00147.k
// func (suite *KTestSuite) TestK00147() {
// 	suite.EqualTestCase(147)
// }

// file://./../testsuite/00148.k
// func (suite *KTestSuite) TestK00148() {
// 	suite.EqualTestCase(148)
// }

// file://./../testsuite/00149.k
// func (suite *KTestSuite) TestK00149() {
// 	suite.EqualTestCase(149)
// }

// file://./../testsuite/00150.k
// func (suite *KTestSuite) TestK00150() {
// 	suite.EqualTestCase(150)
// }

// file://./../testsuite/00151.k
// func (suite *KTestSuite) TestK00151() {
// 	suite.EqualTestCase(151)
// }

// file://./../testsuite/00152.k
// func (suite *KTestSuite) TestK00152() {
// 	suite.EqualTestCase(152)
// }

// file://./../testsuite/00153.k
// func (suite *KTestSuite) TestK00153() {
// 	suite.EqualTestCase(153)
// }

// file://./../testsuite/00154.k
// func (suite *KTestSuite) TestK00154() {
// 	suite.EqualTestCase(154)
// }

// file://./../testsuite/00155.k
// func (suite *KTestSuite) TestK00155() {
// 	suite.EqualTestCase(155)
// }

// file://./../testsuite/00156.k
// func (suite *KTestSuite) TestK00156() {
// 	suite.EqualTestCase(156)
// }

// file://./../testsuite/00157.k
// func (suite *KTestSuite) TestK00157() {
// 	suite.EqualTestCase(157)
// }

// file://./../testsuite/00158.k
// func (suite *KTestSuite) TestK00158() {
// 	suite.EqualTestCase(158)
// }

// file://./../testsuite/00159.k
// func (suite *KTestSuite) TestK00159() {
// 	suite.EqualTestCase(159)
// }

// file://./../testsuite/00160.k
// func (suite *KTestSuite) TestK00160() {
// 	suite.EqualTestCase(160)
// }

// file://./../testsuite/00161.k
// func (suite *KTestSuite) TestK00161() {
// 	suite.EqualTestCase(161)
// }

// file://./../testsuite/00162.k
// func (suite *KTestSuite) TestK00162() {
// 	suite.EqualTestCase(162)
// }

// file://./../testsuite/00163.k
// func (suite *KTestSuite) TestK00163() {
// 	suite.EqualTestCase(163)
// }

// file://./../testsuite/00164.k
// func (suite *KTestSuite) TestK00164() {
// 	suite.EqualTestCase(164)
// }

// file://./../testsuite/00165.k
// func (suite *KTestSuite) TestK00165() {
// 	suite.EqualTestCase(165)
// }

// file://./../testsuite/00166.k
// func (suite *KTestSuite) TestK00166() {
// 	suite.EqualTestCase(166)
// }

// file://./../testsuite/00167.k
// func (suite *KTestSuite) TestK00167() {
// 	suite.EqualTestCase(167)
// }

// file://./../testsuite/00168.k
// func (suite *KTestSuite) TestK00168() {
// 	suite.EqualTestCase(168)
// }

// file://./../testsuite/00169.k
// func (suite *KTestSuite) TestK00169() {
// 	suite.EqualTestCase(169)
// }

// file://./../testsuite/00170.k
// func (suite *KTestSuite) TestK00170() {
// 	suite.EqualTestCase(170)
// }

// file://./../testsuite/00171.k
// func (suite *KTestSuite) TestK00171() {
// 	suite.EqualTestCase(171)
// }

// file://./../testsuite/00172.k
// func (suite *KTestSuite) TestK00172() {
// 	suite.EqualTestCase(172)
// }

// file://./../testsuite/00173.k
// func (suite *KTestSuite) TestK00173() {
// 	suite.EqualTestCase(173)
// }

// file://./../testsuite/00174.k
// func (suite *KTestSuite) TestK00174() {
// 	suite.EqualTestCase(174)
// }

// file://./../testsuite/00175.k
// func (suite *KTestSuite) TestK00175() {
// 	suite.EqualTestCase(175)
// }

// file://./../testsuite/00176.k
// func (suite *KTestSuite) TestK00176() {
// 	suite.EqualTestCase(176)
// }

// file://./../testsuite/00177.k
// func (suite *KTestSuite) TestK00177() {
// 	suite.EqualTestCase(177)
// }

// file://./../testsuite/00178.k
// func (suite *KTestSuite) TestK00178() {
// 	suite.EqualTestCase(178)
// }

// file://./../testsuite/00179.k
func (suite *KTestSuite) TestK00179() {
	suite.EqualTestCase(179)
}

// file://./../testsuite/00180.k
// func (suite *KTestSuite) TestK00180() {
// 	suite.EqualTestCase(180)
// }

// file://./../testsuite/00181.k
// func (suite *KTestSuite) TestK00181() {
// 	suite.EqualTestCase(181)
// }

// file://./../testsuite/00182.k
// func (suite *KTestSuite) TestK00182() {
// 	suite.EqualTestCase(182)
// }

// file://./../testsuite/00183.k
// func (suite *KTestSuite) TestK00183() {
// 	suite.EqualTestCase(183)
// }

// file://./../testsuite/00184.k
// func (suite *KTestSuite) TestK00184() {
// 	suite.EqualTestCase(184)
// }

// file://./../testsuite/00185.k
// func (suite *KTestSuite) TestK00185() {
// 	suite.EqualTestCase(185)
// }

// file://./../testsuite/00186.k
// func (suite *KTestSuite) TestK00186() {
// 	suite.EqualTestCase(186)
// }

// file://./../testsuite/00187.k
// func (suite *KTestSuite) TestK00187() {
// 	suite.EqualTestCase(187)
// }

// file://./../testsuite/00188.k
// func (suite *KTestSuite) TestK00188() {
// 	suite.EqualTestCase(188)
// }

// file://./../testsuite/00189.k
// func (suite *KTestSuite) TestK00189() {
// 	suite.EqualTestCase(189)
// }

// file://./../testsuite/00190.k
// func (suite *KTestSuite) TestK00190() {
// 	suite.EqualTestCase(190)
// }

// file://./../testsuite/00191.k
// func (suite *KTestSuite) TestK00191() {
// 	suite.EqualTestCase(191)
// }

// file://./../testsuite/00192.k
// func (suite *KTestSuite) TestK00192() {
// 	suite.EqualTestCase(192)
// }

// file://./../testsuite/00193.k
// func (suite *KTestSuite) TestK00193() {
// 	suite.EqualTestCase(193)
// }

// file://./../testsuite/00194.k
// func (suite *KTestSuite) TestK00194() {
// 	suite.EqualTestCase(194)
// }

// file://./../testsuite/00195.k
// func (suite *KTestSuite) TestK00195() {
// 	suite.EqualTestCase(195)
// }

// file://./../testsuite/00196.k
// func (suite *KTestSuite) TestK00196() {
// 	suite.EqualTestCase(196)
// }

// file://./../testsuite/00197.k
// func (suite *KTestSuite) TestK00197() {
// 	suite.EqualTestCase(197)
// }

// file://./../testsuite/00198.k
// func (suite *KTestSuite) TestK00198() {
// 	suite.EqualTestCase(198)
// }

// file://./../testsuite/00199.k
// func (suite *KTestSuite) TestK00199() {
// 	suite.EqualTestCase(199)
// }

// file://./../testsuite/00200.k
// func (suite *KTestSuite) TestK00200() {
// 	suite.EqualTestCase(200)
// }

// file://./../testsuite/00201.k
// func (suite *KTestSuite) TestK00201() {
// 	suite.EqualTestCase(201)
// }

// file://./../testsuite/00202.k
// func (suite *KTestSuite) TestK00202() {
// 	suite.EqualTestCase(202)
// }

// file://./../testsuite/00203.k
// func (suite *KTestSuite) TestK00203() {
// 	suite.EqualTestCase(203)
// }

// file://./../testsuite/00204.k
// func (suite *KTestSuite) TestK00204() {
// 	suite.EqualTestCase(204)
// }

// file://./../testsuite/00205.k
// func (suite *KTestSuite) TestK00205() {
// 	suite.EqualTestCase(205)
// }

// file://./../testsuite/00206.k
// func (suite *KTestSuite) TestK00206() {
// 	suite.EqualTestCase(206)
// }

// file://./../testsuite/00207.k
// func (suite *KTestSuite) TestK00207() {
// 	suite.EqualTestCase(207)
// }

// file://./../testsuite/00208.k
// func (suite *KTestSuite) TestK00208() {
// 	suite.EqualTestCase(208)
// }

// file://./../testsuite/00209.k
// func (suite *KTestSuite) TestK00209() {
// 	suite.EqualTestCase(209)
// }

// file://./../testsuite/00210.k
// func (suite *KTestSuite) TestK00210() {
// 	suite.EqualTestCase(210)
// }

// file://./../testsuite/00211.k
// func (suite *KTestSuite) TestK00211() {
// 	suite.EqualTestCase(211)
// }

// file://./../testsuite/00212.k
// func (suite *KTestSuite) TestK00212() {
// 	suite.EqualTestCase(212)
// }

// file://./../testsuite/00213.k
// func (suite *KTestSuite) TestK00213() {
// 	suite.EqualTestCase(213)
// }

// file://./../testsuite/00214.k
// func (suite *KTestSuite) TestK00214() {
// 	suite.EqualTestCase(214)
// }

// file://./../testsuite/00215.k
// func (suite *KTestSuite) TestK00215() {
// 	suite.EqualTestCase(215)
// }

// file://./../testsuite/00216.k
// func (suite *KTestSuite) TestK00216() {
// 	suite.EqualTestCase(216)
// }

// file://./../testsuite/00217.k
// func (suite *KTestSuite) TestK00217() {
// 	suite.EqualTestCase(217)
// }

// file://./../testsuite/00218.k
// func (suite *KTestSuite) TestK00218() {
// 	suite.EqualTestCase(218)
// }

// file://./../testsuite/00219.k
// func (suite *KTestSuite) TestK00219() {
// 	suite.EqualTestCase(219)
// }

// file://./../testsuite/00220.k
// func (suite *KTestSuite) TestK00220() {
// 	suite.EqualTestCase(220)
// }

func TestKTestSuite(t *t.T) {
	suite.Run(t, new(KTestSuite))
}
