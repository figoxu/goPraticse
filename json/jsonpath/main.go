package main

import (
	"encoding/json"
	"github.com/figoxu/utee"
	"github.com/oliveagle/jsonpath"
	"log"
)

var jsonVal = `{"status":"1","info":"OK","infocode":"10000","count":"1","route":{"origin":"116.481028,39.989643","destination":"116.434446,39.90816","taxi_cost":"35.004099999999994","paths":[{"distance":"12567","duration":"1961","strategy":"距离最短","tolls":"0","toll_distance":"0","steps":[{"instruction":"向西南行驶44米右转进入主路","orientation":"西南","distance":"44","tolls":"0","toll_distance":"0","toll_road":[],"duration":"13","polyline":"116.481216,39.989536;116.481003,39.989311;116.480949,39.989265;116.480896,39.989212","action":"右转","assistant_action":"进入主路","tmcs":[{"lcode":[],"distance":"31","status":"未知","polyline":"116.481216,39.989536;116.481003,39.989311"},{"lcode":[],"distance":"6","status":"未知","polyline":"116.481003,39.989311;116.480949,39.989265"},{"lcode":[],"distance":"7","status":"未知","polyline":"116.480949,39.989265;116.480896,39.989212"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"朝阳区","adcode":"110105"}]}]},{"instruction":"沿阜荣街向西北行驶112米左转进入主路","orientation":"西北","road":"阜荣街","distance":"112","tolls":"0","toll_distance":"0","toll_road":[],"duration":"40","polyline":"116.480835,39.989155;116.479843,39.989815","action":"左转","assistant_action":"进入主路","tmcs":[{"lcode":[],"distance":"112","status":"畅通","polyline":"116.480835,39.989155;116.479843,39.989815"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"朝阳区","adcode":"110105"}]}]},{"instruction":"沿阜通东大街向西南行驶267米左转进入主路","orientation":"西南","road":"阜通东大街","distance":"267","tolls":"0","toll_distance":"0","toll_road":[],"duration":"60","polyline":"116.47963,39.98983;116.479164,39.989403;116.479118,39.989368;116.479095,39.989338;116.478752,39.989033;116.477592,39.98801","action":"左转","assistant_action":"进入主路","tmcs":[{"lcode":[],"distance":"61","status":"畅通","polyline":"116.47963,39.98983;116.479164,39.989403"},{"lcode":[],"distance":"10","status":"畅通","polyline":"116.479164,39.989403;116.479118,39.989368;116.479095,39.989338"},{"lcode":[],"distance":"44","status":"畅通","polyline":"116.479095,39.989338;116.478752,39.989033"},{"lcode":[],"distance":"152","status":"畅通","polyline":"116.478752,39.989033;116.477592,39.98801"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"朝阳区","adcode":"110105"}]}]},{"instruction":"沿广顺南大街向东南行驶563米右转","orientation":"东南","road":"广顺南大街","distance":"563","tolls":"0","toll_distance":"0","toll_road":[],"duration":"163","polyline":"116.4776,39.987698;116.478088,39.987385;116.478752,39.986923;116.479225,39.986599;116.479416,39.986469;116.480453,39.985737;116.480827,39.985477;116.481148,39.985271;116.481247,39.985199;116.481575,39.984985;116.482254,39.984535;116.48259,39.984333","action":"右转","assistant_action":[],"tmcs":[{"lcode":[],"distance":"131","status":"畅通","polyline":"116.4776,39.987698;116.478088,39.987385;116.478752,39.986923"},{"lcode":[],"distance":"54","status":"畅通","polyline":"116.478752,39.986923;116.479225,39.986599"},{"lcode":[],"distance":"21","status":"畅通","polyline":"116.479225,39.986599;116.479416,39.986469"},{"lcode":[],"distance":"121","status":"畅通","polyline":"116.479416,39.986469;116.480453,39.985737"},{"lcode":[],"distance":"42","status":"畅通","polyline":"116.480453,39.985737;116.480827,39.985477"},{"lcode":[],"distance":"46","status":"畅通","polyline":"116.480827,39.985477;116.481148,39.985271;116.481247,39.985199"},{"lcode":[],"distance":"36","status":"畅通","polyline":"116.481247,39.985199;116.481575,39.984985"},{"lcode":[],"distance":"77","status":"畅通","polyline":"116.481575,39.984985;116.482254,39.984535"},{"lcode":[],"distance":"35","status":"畅通","polyline":"116.482254,39.984535;116.48259,39.984333"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"朝阳区","adcode":"110105"}]}]},{"instruction":"沿京密路途径香河园路向西南行驶4.9千米左转","orientation":"西南","road":"京密路","distance":"4860","tolls":"0","toll_distance":"0","toll_road":[],"duration":"682","polyline":"116.48259,39.984333;116.482498,39.984257;116.481903,39.983723;116.481659,39.98349;116.481224,39.983082;116.479645,39.981667;116.478668,39.980793;116.478088,39.980267;116.47802,39.980202;116.477982,39.980164;116.477211,39.979435;116.477051,39.979294;116.477005,39.979244;116.476852,39.979111;116.476707,39.978992;116.476044,39.978382;116.47509,39.977531;116.475029,39.977474;116.474304,39.976814;116.473877,39.976421;116.473282,39.975887;116.472763,39.975426;116.472382,39.975105;116.471848,39.974628;116.471313,39.974148;116.469536,39.972515;116.469086,39.972118;116.468277,39.971397;116.467865,39.971024;116.466988,39.970222;116.466515,39.969799;116.46537,39.968746;116.464951,39.968369;116.464836,39.968258;116.464775,39.968208;116.464462,39.967926;116.462807,39.966431;116.462746,39.966377;116.460793,39.964596;116.460114,39.963982;116.459984,39.963852;116.459404,39.963341;116.459129,39.963089;116.458076,39.962132;116.457954,39.962021;116.457909,39.961983;116.457359,39.961475;116.457306,39.961422;116.456764,39.960968;116.456657,39.960876;116.456047,39.960472;116.455536,39.960209;116.455078,39.959991;116.454765,39.959839;116.45459,39.95974;116.454437,39.959641;116.454178,39.959438;116.453568,39.958889;116.453384,39.958721;116.453041,39.958408;116.452881,39.95826;116.452232,39.957672;116.451797,39.957275;116.450981,39.956547;116.450882,39.956459;116.450859,39.95644;116.450676,39.95628;116.450333,39.95599;116.449989,39.955688;116.449707,39.955437;116.449135,39.954914;116.448853,39.954639;116.448311,39.95417;116.447609,39.953632;116.447571,39.953606;116.447014,39.953194;116.445831,39.952305;116.445297,39.951923;116.444717,39.951519","action":"左转","assistant_action":[],"tmcs":[{"lcode":[],"distance":"11","status":"畅通","polyline":"116.48259,39.984333;116.482498,39.984257"},{"lcode":[],"distance":"376","status":"畅通","polyline":"116.482498,39.984257;116.481903,39.983723;116.481659,39.98349;116.481224,39.983082;116.479645,39.981667"},{"lcode":[],"distance":"213","status":"畅通","polyline":"116.479645,39.981667;116.478668,39.980793;116.478088,39.980267;116.47802,39.980202"},{"lcode":[],"distance":"5","status":"畅通","polyline":"116.47802,39.980202;116.477982,39.980164"},{"lcode":[],"distance":"105","status":"畅通","polyline":"116.477982,39.980164;116.477211,39.979435"},{"lcode":[],"distance":"46","status":"畅通","polyline":"116.477211,39.979435;116.477051,39.979294;116.477005,39.979244;116.476852,39.979111"},{"lcode":[],"distance":"17","status":"畅通","polyline":"116.476852,39.979111;116.476707,39.978992"},{"lcode":[],"distance":"88","status":"畅通","polyline":"116.476707,39.978992;116.476044,39.978382"},{"lcode":[],"distance":"132","status":"畅通","polyline":"116.476044,39.978382;116.47509,39.977531;116.475029,39.977474"},{"lcode":[],"distance":"96","status":"畅通","polyline":"116.475029,39.977474;116.474304,39.976814"},{"lcode":[],"distance":"56","status":"畅通","polyline":"116.474304,39.976814;116.473877,39.976421"},{"lcode":[],"distance":"78","status":"畅通","polyline":"116.473877,39.976421;116.473282,39.975887"},{"lcode":[],"distance":"67","status":"畅通","polyline":"116.473282,39.975887;116.472763,39.975426"},{"lcode":[],"distance":"47","status":"畅通","polyline":"116.472763,39.975426;116.472382,39.975105"},{"lcode":[],"distance":"140","status":"畅通","polyline":"116.472382,39.975105;116.471848,39.974628;116.471313,39.974148"},{"lcode":[],"distance":"295","status":"畅通","polyline":"116.471313,39.974148;116.469536,39.972515;116.469086,39.972118"},{"lcode":[],"distance":"105","status":"畅通","polyline":"116.469086,39.972118;116.468277,39.971397"},{"lcode":[],"distance":"232","status":"畅通","polyline":"116.468277,39.971397;116.467865,39.971024;116.466988,39.970222;116.466515,39.969799"},{"lcode":[],"distance":"152","status":"畅通","polyline":"116.466515,39.969799;116.46537,39.968746"},{"lcode":[],"distance":"55","status":"畅通","polyline":"116.46537,39.968746;116.464951,39.968369"},{"lcode":[],"distance":"289","status":"畅通","polyline":"116.464951,39.968369;116.464836,39.968258;116.464775,39.968208;116.464462,39.967926;116.462807,39.966431;116.462746,39.966377"},{"lcode":[],"distance":"258","status":"缓行","polyline":"116.462746,39.966377;116.460793,39.964596"},{"lcode":[],"distance":"107","status":"缓行","polyline":"116.460793,39.964596;116.460114,39.963982;116.459984,39.963852"},{"lcode":[],"distance":"76","status":"缓行","polyline":"116.459984,39.963852;116.459404,39.963341"},{"lcode":[],"distance":"175","status":"缓行","polyline":"116.459404,39.963341;116.459129,39.963089;116.458076,39.962132"},{"lcode":[],"distance":"21","status":"畅通","polyline":"116.458076,39.962132;116.457954,39.962021;116.457909,39.961983"},{"lcode":[],"distance":"80","status":"畅通","polyline":"116.457909,39.961983;116.457359,39.961475;116.457306,39.961422"},{"lcode":[],"distance":"82","status":"畅通","polyline":"116.457306,39.961422;116.456764,39.960968;116.456657,39.960876"},{"lcode":[],"distance":"121","status":"畅通","polyline":"116.456657,39.960876;116.456047,39.960472;116.455536,39.960209"},{"lcode":[],"distance":"144","status":"畅通","polyline":"116.455536,39.960209;116.455078,39.959991;116.454765,39.959839;116.45459,39.95974;116.454437,39.959641;116.454178,39.959438"},{"lcode":[],"distance":"257","status":"畅通","polyline":"116.454178,39.959438;116.453568,39.958889;116.453384,39.958721;116.453041,39.958408;116.452881,39.95826;116.452232,39.957672"},{"lcode":[],"distance":"164","status":"畅通","polyline":"116.452232,39.957672;116.451797,39.957275;116.450981,39.956547"},{"lcode":[],"distance":"15","status":"畅通","polyline":"116.450981,39.956547;116.450882,39.956459;116.450859,39.95644"},{"lcode":[],"distance":"24","status":"畅通","polyline":"116.450859,39.95644;116.450676,39.95628"},{"lcode":[],"distance":"87","status":"缓行","polyline":"116.450676,39.95628;116.450333,39.95599;116.449989,39.955688"},{"lcode":[],"distance":"37","status":"畅通","polyline":"116.449989,39.955688;116.449707,39.955437"},{"lcode":[],"distance":"75","status":"畅通","polyline":"116.449707,39.955437;116.449135,39.954914"},{"lcode":[],"distance":"197","status":"畅通","polyline":"116.449135,39.954914;116.448853,39.954639;116.448311,39.95417;116.447609,39.953632;116.447571,39.953606"},{"lcode":[],"distance":"65","status":"畅通","polyline":"116.447571,39.953606;116.447014,39.953194"},{"lcode":[],"distance":"141","status":"畅通","polyline":"116.447014,39.953194;116.445831,39.952305"},{"lcode":[],"distance":"62","status":"畅通","polyline":"116.445831,39.952305;116.445297,39.951923"},{"lcode":[],"distance":"67","status":"畅通","polyline":"116.445297,39.951923;116.444717,39.951519"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"朝阳区","adcode":"110105"}]}]},{"instruction":"沿新东路向东南行驶154米右转","orientation":"东南","road":"新东路","distance":"154","tolls":"0","toll_distance":"0","toll_road":[],"duration":"34","polyline":"116.44474,39.951385;116.444878,39.95121;116.444962,39.951138;116.445274,39.950954;116.446053,39.950439","action":"右转","assistant_action":[],"tmcs":[{"lcode":[],"distance":"154","status":"缓行","polyline":"116.44474,39.951385;116.444878,39.95121;116.444962,39.951138;116.445274,39.950954;116.446053,39.950439"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"朝阳区","adcode":"110105"}]}]},{"instruction":"沿东直门外斜街向西南行驶1.3千米右转","orientation":"西南","road":"东直门外斜街","distance":"1251","tolls":"0","toll_distance":"0","toll_road":[],"duration":"259","polyline":"116.446053,39.950439;116.4459,39.950218;116.445763,39.950043;116.445274,39.9496;116.444649,39.949028;116.444336,39.94875;116.443314,39.947861;116.443085,39.947662;116.443001,39.947601;116.442528,39.947201;116.44223,39.946964;116.441521,39.946346;116.44146,39.946289;116.440628,39.94553;116.440506,39.945412;116.440376,39.945278;116.440231,39.945148;116.439713,39.944702;116.438881,39.943928;116.438744,39.943775;116.438683,39.943657;116.438583,39.943432;116.438545,39.943077;116.438545,39.942894;116.43856,39.942463;116.438599,39.941498;116.438614,39.941193","action":"右转","assistant_action":[],"tmcs":[{"lcode":[],"distance":"115","status":"畅通","polyline":"116.446053,39.950439;116.4459,39.950218;116.445763,39.950043;116.445274,39.9496"},{"lcode":[],"distance":"82","status":"畅通","polyline":"116.445274,39.9496;116.444649,39.949028"},{"lcode":[],"distance":"40","status":"畅通","polyline":"116.444649,39.949028;116.444336,39.94875"},{"lcode":[],"distance":"132","status":"畅通","polyline":"116.444336,39.94875;116.443314,39.947861"},{"lcode":[],"distance":"39","status":"畅通","polyline":"116.443314,39.947861;116.443085,39.947662;116.443001,39.947601"},{"lcode":[],"distance":"60","status":"畅通","polyline":"116.443001,39.947601;116.442528,39.947201"},{"lcode":[],"distance":"35","status":"畅通","polyline":"116.442528,39.947201;116.44223,39.946964"},{"lcode":[],"distance":"92","status":"畅通","polyline":"116.44223,39.946964;116.441521,39.946346"},{"lcode":[],"distance":"118","status":"畅通","polyline":"116.441521,39.946346;116.44146,39.946289;116.440628,39.94553"},{"lcode":[],"distance":"34","status":"畅通","polyline":"116.440628,39.94553;116.440506,39.945412;116.440376,39.945278"},{"lcode":[],"distance":"19","status":"畅通","polyline":"116.440376,39.945278;116.440231,39.945148"},{"lcode":[],"distance":"66","status":"畅通","polyline":"116.440231,39.945148;116.439713,39.944702"},{"lcode":[],"distance":"173","status":"缓行","polyline":"116.439713,39.944702;116.438881,39.943928;116.438744,39.943775;116.438683,39.943657;116.438583,39.943432"},{"lcode":[],"distance":"38","status":"畅通","polyline":"116.438583,39.943432;116.438545,39.943077"},{"lcode":[],"distance":"20","status":"拥堵","polyline":"116.438545,39.943077;116.438545,39.942894"},{"lcode":[],"distance":"48","status":"拥堵","polyline":"116.438545,39.942894;116.43856,39.942463"},{"lcode":[],"distance":"107","status":"缓行","polyline":"116.43856,39.942463;116.438599,39.941498"},{"lcode":[],"distance":"33","status":"缓行","polyline":"116.438599,39.941498;116.438614,39.941193"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"朝阳区","adcode":"110105"},{"name":"东城区","adcode":"110101"}]}]},{"instruction":"沿东直门外大街途径东直门桥向西行驶439米靠左","orientation":"西","road":"东直门外大街","distance":"439","tolls":"0","toll_distance":"0","toll_road":[],"duration":"61","polyline":"116.438614,39.941193;116.437614,39.941158;116.4375,39.941158;116.437439,39.941158;116.437088,39.941158;116.436401,39.941147;116.436356,39.941147;116.435814,39.941147;116.435402,39.941147;116.435379,39.941147;116.435104,39.941154;116.434944,39.941158;116.434761,39.941166;116.434578,39.941227;116.434258,39.941372;116.434128,39.94141;116.434059,39.941406;116.433517,39.941395;116.43351,39.941395","action":"靠左","assistant_action":[],"tmcs":[{"lcode":[],"distance":"84","status":"畅通","polyline":"116.438614,39.941193;116.437614,39.941158"},{"lcode":[],"distance":"15","status":"畅通","polyline":"116.437614,39.941158;116.4375,39.941158;116.437439,39.941158"},{"lcode":[],"distance":"30","status":"畅通","polyline":"116.437439,39.941158;116.437088,39.941158"},{"lcode":[],"distance":"61","status":"畅通","polyline":"116.437088,39.941158;116.436401,39.941147;116.436356,39.941147"},{"lcode":[],"distance":"136","status":"畅通","polyline":"116.436356,39.941147;116.435814,39.941147;116.435402,39.941147;116.435379,39.941147;116.435104,39.941154;116.434944,39.941158;116.434761,39.941166"},{"lcode":[],"distance":"60","status":"畅通","polyline":"116.434761,39.941166;116.434578,39.941227;116.434258,39.941372;116.434128,39.94141"},{"lcode":[],"distance":"53","status":"缓行","polyline":"116.434128,39.94141;116.434059,39.941406;116.433517,39.941395;116.43351,39.941395"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"东城区","adcode":"110101"}]}]},{"instruction":"沿东直门桥向西行驶31米靠左","orientation":"西","road":"东直门桥","distance":"31","tolls":"0","toll_distance":"0","toll_road":[],"duration":"6","polyline":"116.43351,39.941395;116.433235,39.941326;116.433182,39.941265","action":"靠左","assistant_action":[],"tmcs":[{"lcode":[],"distance":"24","status":"缓行","polyline":"116.43351,39.941395;116.433235,39.941326"},{"lcode":[],"distance":"7","status":"缓行","polyline":"116.433235,39.941326;116.433182,39.941265"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"东城区","adcode":"110101"}]}]},{"instruction":"沿东直门桥向南行驶49米向右前方行驶进入匝道","orientation":"南","road":"东直门桥","distance":"49","tolls":"0","toll_distance":"0","toll_road":[],"duration":"12","polyline":"116.433182,39.941265;116.433144,39.941154;116.433144,39.941074;116.433159,39.941025;116.43322,39.940887;116.433281,39.940849","action":"向右前方行驶","assistant_action":"进入匝道","tmcs":[{"lcode":[],"distance":"43","status":"缓行","polyline":"116.433182,39.941265;116.433144,39.941154;116.433144,39.941074;116.433159,39.941025;116.43322,39.940887"},{"lcode":[],"distance":"6","status":"缓行","polyline":"116.43322,39.940887;116.433281,39.940849"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"东城区","adcode":"110101"}]}]},{"instruction":"沿东直门桥途径东直门南大街向南行驶527米向左前方行驶进入主路","orientation":"南","road":"东直门桥","distance":"527","tolls":"0","toll_distance":"0","toll_road":[],"duration":"106","polyline":"116.433281,39.940849;116.433327,39.940701;116.433357,39.940575;116.433372,39.940315;116.433372,39.940224;116.433357,39.94009;116.433273,39.939678;116.433266,39.93959;116.433266,39.939533;116.433289,39.939484;116.433487,39.939159;116.433556,39.938992;116.433655,39.938824;116.433701,39.937702;116.433708,39.937454;116.433716,39.937157;116.433739,39.936707;116.433762,39.936321;116.433769,39.936131","action":"向左前方行驶","assistant_action":"进入主路","tmcs":[{"lcode":[],"distance":"31","status":"缓行","polyline":"116.433281,39.940849;116.433327,39.940701;116.433357,39.940575"},{"lcode":[],"distance":"109","status":"拥堵","polyline":"116.433357,39.940575;116.433372,39.940315;116.433372,39.940224;116.433357,39.94009;116.433273,39.939678;116.433266,39.93959"},{"lcode":[],"distance":"91","status":"缓行","polyline":"116.433266,39.93959;116.433266,39.939533;116.433289,39.939484;116.433487,39.939159;116.433556,39.938992;116.433655,39.938824"},{"lcode":[],"distance":"124","status":"缓行","polyline":"116.433655,39.938824;116.433701,39.937702"},{"lcode":[],"distance":"27","status":"缓行","polyline":"116.433701,39.937702;116.433708,39.937454"},{"lcode":[],"distance":"32","status":"缓行","polyline":"116.433708,39.937454;116.433716,39.937157"},{"lcode":[],"distance":"50","status":"缓行","polyline":"116.433716,39.937157;116.433739,39.936707"},{"lcode":[],"distance":"42","status":"缓行","polyline":"116.433739,39.936707;116.433762,39.936321"},{"lcode":[],"distance":"21","status":"畅通","polyline":"116.433762,39.936321;116.433769,39.936131"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"东城区","adcode":"110101"}]}]},{"instruction":"沿东二环入口途径东二环向南行驶3.4千米右转","orientation":"南","road":"东二环入口","distance":"3396","tolls":"0","toll_distance":"0","toll_road":[],"duration":"318","polyline":"116.433769,39.936131;116.433937,39.935814;116.433975,39.935299;116.434044,39.934086;116.43409,39.93232;116.434143,39.931446;116.434143,39.931133;116.434151,39.931068;116.434166,39.93037;116.434212,39.929539;116.43425,39.928741;116.434265,39.928299;116.434341,39.926754;116.434448,39.924316;116.434525,39.922337;116.43454,39.921963;116.434563,39.921654;116.434586,39.920837;116.434639,39.919716;116.434723,39.917976;116.434723,39.917671;116.434807,39.91597;116.43483,39.915417;116.434837,39.91526;116.434883,39.914467;116.434921,39.913979;116.434998,39.913506;116.435028,39.91333;116.435074,39.913048;116.435333,39.912014;116.435501,39.911263;116.435555,39.910969;116.435593,39.910671;116.435608,39.910358;116.435646,39.909748;116.435753,39.907406;116.435822,39.906105;116.435837,39.905937;116.435905,39.90554","action":"右转","assistant_action":[],"tmcs":[{"lcode":[],"distance":"38","status":"缓行","polyline":"116.433769,39.936131;116.433937,39.935814"},{"lcode":[],"distance":"485","status":"畅通","polyline":"116.433937,39.935814;116.433975,39.935299;116.434044,39.934086;116.43409,39.93232;116.434143,39.931446"},{"lcode":[],"distance":"41","status":"缓行","polyline":"116.434143,39.931446;116.434143,39.931133;116.434151,39.931068"},{"lcode":[],"distance":"77","status":"缓行","polyline":"116.434151,39.931068;116.434166,39.93037"},{"lcode":[],"distance":"93","status":"缓行","polyline":"116.434166,39.93037;116.434212,39.929539"},{"lcode":[],"distance":"88","status":"缓行","polyline":"116.434212,39.929539;116.43425,39.928741"},{"lcode":[],"distance":"48","status":"缓行","polyline":"116.43425,39.928741;116.434265,39.928299"},{"lcode":[],"distance":"171","status":"缓行","polyline":"116.434265,39.928299;116.434341,39.926754"},{"lcode":[],"distance":"490","status":"缓行","polyline":"116.434341,39.926754;116.434448,39.924316;116.434525,39.922337"},{"lcode":[],"distance":"76","status":"缓行","polyline":"116.434525,39.922337;116.43454,39.921963;116.434563,39.921654"},{"lcode":[],"distance":"91","status":"缓行","polyline":"116.434563,39.921654;116.434586,39.920837"},{"lcode":[],"distance":"123","status":"缓行","polyline":"116.434586,39.920837;116.434639,39.919716"},{"lcode":[],"distance":"193","status":"缓行","polyline":"116.434639,39.919716;116.434723,39.917976"},{"lcode":[],"distance":"34","status":"畅通","polyline":"116.434723,39.917976;116.434723,39.917671"},{"lcode":[],"distance":"189","status":"畅通","polyline":"116.434723,39.917671;116.434807,39.91597"},{"lcode":[],"distance":"78","status":"畅通","polyline":"116.434807,39.91597;116.43483,39.915417;116.434837,39.91526"},{"lcode":[],"distance":"142","status":"畅通","polyline":"116.434837,39.91526;116.434883,39.914467;116.434921,39.913979"},{"lcode":[],"distance":"52","status":"畅通","polyline":"116.434921,39.913979;116.434998,39.913506"},{"lcode":[],"distance":"51","status":"畅通","polyline":"116.434998,39.913506;116.435028,39.91333;116.435074,39.913048"},{"lcode":[],"distance":"235","status":"畅通","polyline":"116.435074,39.913048;116.435333,39.912014;116.435501,39.911263;116.435555,39.910969"},{"lcode":[],"distance":"135","status":"畅通","polyline":"116.435555,39.910969;116.435593,39.910671;116.435608,39.910358;116.435646,39.909748"},{"lcode":[],"distance":"260","status":"畅通","polyline":"116.435646,39.909748;116.435753,39.907406"},{"lcode":[],"distance":"144","status":"畅通","polyline":"116.435753,39.907406;116.435822,39.906105"},{"lcode":[],"distance":"18","status":"畅通","polyline":"116.435822,39.906105;116.435837,39.905937"},{"lcode":[],"distance":"44","status":"畅通","polyline":"116.435837,39.905937;116.435905,39.90554"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"东城区","adcode":"110101"}]}]},{"instruction":"沿东二环出口途径北京站东街向西行驶321米右转","orientation":"西","road":"东二环出口","distance":"321","tolls":"0","toll_distance":"0","toll_road":[],"duration":"63","polyline":"116.435905,39.90554;116.435875,39.905445;116.435814,39.90535;116.435799,39.905323;116.435753,39.905285;116.435699,39.905258;116.435608,39.905224;116.435486,39.905197;116.434807,39.905174;116.43354,39.90517;116.433334,39.905174;116.433189,39.90517;116.433105,39.90517;116.433006,39.905163;116.432381,39.905132","action":"右转","assistant_action":[],"tmcs":[{"lcode":[],"distance":"38","status":"畅通","polyline":"116.435905,39.90554;116.435875,39.905445;116.435814,39.90535;116.435799,39.905323;116.435753,39.905285;116.435699,39.905258"},{"lcode":[],"distance":"19","status":"畅通","polyline":"116.435699,39.905258;116.435608,39.905224;116.435486,39.905197"},{"lcode":[],"distance":"58","status":"畅通","polyline":"116.435486,39.905197;116.434807,39.905174"},{"lcode":[],"distance":"108","status":"畅通","polyline":"116.434807,39.905174;116.43354,39.90517"},{"lcode":[],"distance":"17","status":"畅通","polyline":"116.43354,39.90517;116.433334,39.905174"},{"lcode":[],"distance":"81","status":"畅通","polyline":"116.433334,39.905174;116.433189,39.90517;116.433105,39.90517;116.433006,39.905163;116.432381,39.905132"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"东城区","adcode":"110101"}]}]},{"instruction":"沿大羊毛胡同向北行驶365米右转进入辅路","orientation":"北","road":"大羊毛胡同","distance":"365","tolls":"0","toll_distance":"0","toll_road":[],"duration":"74","polyline":"116.432381,39.905132;116.432312,39.905876;116.432312,39.906265;116.432312,39.906651;116.432304,39.906876;116.432266,39.907867;116.432259,39.90844","action":"右转","assistant_action":"进入辅路","tmcs":[{"lcode":[],"distance":"82","status":"畅通","polyline":"116.432381,39.905132;116.432312,39.905876"},{"lcode":[],"distance":"44","status":"畅通","polyline":"116.432312,39.905876;116.432312,39.906265"},{"lcode":[],"distance":"41","status":"畅通","polyline":"116.432312,39.906265;116.432312,39.906651"},{"lcode":[],"distance":"25","status":"畅通","polyline":"116.432312,39.906651;116.432304,39.906876"},{"lcode":[],"distance":"109","status":"畅通","polyline":"116.432304,39.906876;116.432266,39.907867"},{"lcode":[],"distance":"64","status":"畅通","polyline":"116.432266,39.907867;116.432259,39.90844"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"东城区","adcode":"110101"}]}]},{"instruction":"沿建国门内大街向东行驶188米到达目的地","orientation":"东","road":"建国门内大街","distance":"188","tolls":"0","toll_distance":"0","toll_road":[],"duration":"70","polyline":"116.432259,39.90844;116.432495,39.908375;116.433105,39.908379;116.434349,39.908401;116.434448,39.908401","action":[],"assistant_action":"到达目的地","tmcs":[{"lcode":[],"distance":"188","status":"拥堵","polyline":"116.432259,39.90844;116.432495,39.908375;116.433105,39.908379;116.434349,39.908401;116.434448,39.908401"}],"cities":[{"name":"北京城区","citycode":"010","adcode":"110100","districts":[{"name":"东城区","adcode":"110101"}]}]}],"restriction":"0","traffic_lights":"15"}]}}`

func main() {
	parse(jsonVal, "$.route.paths[0].distance")
}

func parse(data, jspath string) string {
	var json_data interface{}
	err := json.Unmarshal([]byte(data), &json_data)
	utee.Chk(err)
	vs, err := jsonpath.JsonPathLookup(json_data, jspath)
	utee.Chk(err)
	log.Println(vs)
	return ""
}
