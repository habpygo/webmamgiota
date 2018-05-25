/*
Copyright DappDevelopment. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

/* Package pathdata contains the paths, addresses,
http:// addresses and seeds
*/
package metadata

var Seed = "THISISTHETESTSENTENCETOEXPERIMENTWITHIOTATANGLEFORPROGRAMMINGUSECASESASWELLASFUN9"
var Address = "TVWZVZZLWSMLXYTFQNVQSAGCQLRRCUXMUDDQWJILNQGOIFKMA9PKBRKORIWOOF9WQLJWGVGTWUXPNNKNYSRBAWUWQC"

var Provider = "https://nodes.testnet.thetangle.org:443"

var MWM int64 = 9 // In the real world set this to 14 or 15!

/*
Min Weight Magnitude (MWM) - the amount of Work that will be carried out in the PoW stage.
this means that a solution to the puzzle is a number with MWM trailing 0's (9's in trytes).
currently MWM is set to 14 on the mainnet and 9 on the testnet.
each increment of MWM is 3 times harder PoW (on average).
*/
