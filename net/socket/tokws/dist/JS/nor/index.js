var swfobject = function () { var D = "undefined", r = "object", S = "Shockwave Flash", W = "ShockwaveFlash.ShockwaveFlash", q = "application/x-shockwave-flash", R = "SWFObjectExprInst", x = "onreadystatechange", O = window, j = document, t = navigator, T = false, U = [h], o = [], N = [], I = [], l, Q, E, B, J = false, a = false, n, G, m = true, M = function () { var aa = typeof j.getElementById != D && typeof j.getElementsByTagName != D && typeof j.createElement != D, ah = t.userAgent.toLowerCase(), Y = t.platform.toLowerCase(), ae = Y ? /win/.test(Y) : /win/.test(ah), ac = Y ? /mac/.test(Y) : /mac/.test(ah), af = /webkit/.test(ah) ? parseFloat(ah.replace(/^.*webkit\/(\d+(\.\d+)?).*$/, "$1")) : false, X = !+"\v1", ag = [0, 0, 0], ab = null; if (typeof t.plugins != D && typeof t.plugins[S] == r) { ab = t.plugins[S].description; if (ab && !(typeof t.mimeTypes != D && t.mimeTypes[q] && !t.mimeTypes[q].enabledPlugin)) { T = true; X = false; ab = ab.replace(/^.*\s+(\S+\s+\S+$)/, "$1"); ag[0] = parseInt(ab.replace(/^(.*)\..*$/, "$1"), 10); ag[1] = parseInt(ab.replace(/^.*\.(.*)\s.*$/, "$1"), 10); ag[2] = /[a-zA-Z]/.test(ab) ? parseInt(ab.replace(/^.*[a-zA-Z]+(.*)$/, "$1"), 10) : 0 } } else { if (typeof O.ActiveXObject != D) { try { var ad = new ActiveXObject(W); if (ad) { ab = ad.GetVariable("$version"); if (ab) { X = true; ab = ab.split(" ")[1].split(","); ag = [parseInt(ab[0], 10), parseInt(ab[1], 10), parseInt(ab[2], 10)] } } } catch (Z) { } } } return { w3: aa, pv: ag, wk: af, ie: X, win: ae, mac: ac } }(), k = function () { if (!M.w3) { return } if ((typeof j.readyState != D && j.readyState == "complete") || (typeof j.readyState == D && (j.getElementsByTagName("body")[0] || j.body))) { f() } if (!J) { if (typeof j.addEventListener != D) { j.addEventListener("DOMContentLoaded", f, false) } if (M.ie && M.win) { j.attachEvent(x, function () { if (j.readyState == "complete") { j.detachEvent(x, arguments.callee); f() } }); if (O == top) { (function () { if (J) { return } try { j.documentElement.doScroll("left") } catch (X) { setTimeout(arguments.callee, 0); return } f() })() } } if (M.wk) { (function () { if (J) { return } if (!/loaded|complete/.test(j.readyState)) { setTimeout(arguments.callee, 0); return } f() })() } s(f) } }(); function f() { if (J) { return } try { var Z = j.getElementsByTagName("body")[0].appendChild(C("span")); Z.parentNode.removeChild(Z) } catch (aa) { return } J = true; var X = U.length; for (var Y = 0; Y < X; Y++) { U[Y]() } } function K(X) { if (J) { X() } else { U[U.length] = X } } function s(Y) { if (typeof O.addEventListener != D) { O.addEventListener("load", Y, false) } else { if (typeof j.addEventListener != D) { j.addEventListener("load", Y, false) } else { if (typeof O.attachEvent != D) { i(O, "onload", Y) } else { if (typeof O.onload == "function") { var X = O.onload; O.onload = function () { X(); Y() } } else { O.onload = Y } } } } } function h() { if (T) { V() } else { H() } } function V() { var X = j.getElementsByTagName("body")[0]; var aa = C(r); aa.setAttribute("type", q); var Z = X.appendChild(aa); if (Z) { var Y = 0; (function () { if (typeof Z.GetVariable != D) { var ab = Z.GetVariable("$version"); if (ab) { ab = ab.split(" ")[1].split(","); M.pv = [parseInt(ab[0], 10), parseInt(ab[1], 10), parseInt(ab[2], 10)] } } else { if (Y < 10) { Y++; setTimeout(arguments.callee, 10); return } } X.removeChild(aa); Z = null; H() })() } else { H() } } function H() { var ag = o.length; if (ag > 0) { for (var af = 0; af < ag; af++) { var Y = o[af].id; var ab = o[af].callbackFn; var aa = { success: false, id: Y }; if (M.pv[0] > 0) { var ae = c(Y); if (ae) { if (F(o[af].swfVersion) && !(M.wk && M.wk < 312)) { w(Y, true); if (ab) { aa.success = true; aa.ref = z(Y); ab(aa) } } else { if (o[af].expressInstall && A()) { var ai = {}; ai.data = o[af].expressInstall; ai.width = ae.getAttribute("width") || "0"; ai.height = ae.getAttribute("height") || "0"; if (ae.getAttribute("class")) { ai.styleclass = ae.getAttribute("class") } if (ae.getAttribute("align")) { ai.align = ae.getAttribute("align") } var ah = {}; var X = ae.getElementsByTagName("param"); var ac = X.length; for (var ad = 0; ad < ac; ad++) { if (X[ad].getAttribute("name").toLowerCase() != "movie") { ah[X[ad].getAttribute("name")] = X[ad].getAttribute("value") } } P(ai, ah, Y, ab) } else { p(ae); if (ab) { ab(aa) } } } } } else { w(Y, true); if (ab) { var Z = z(Y); if (Z && typeof Z.SetVariable != D) { aa.success = true; aa.ref = Z } ab(aa) } } } } } function z(aa) { var X = null; var Y = c(aa); if (Y && Y.nodeName == "OBJECT") { if (typeof Y.SetVariable != D) { X = Y } else { var Z = Y.getElementsByTagName(r)[0]; if (Z) { X = Z } } } return X } function A() { return !a && F("6.0.65") && (M.win || M.mac) && !(M.wk && M.wk < 312) } function P(aa, ab, X, Z) { a = true; E = Z || null; B = { success: false, id: X }; var ae = c(X); if (ae) { if (ae.nodeName == "OBJECT") { l = g(ae); Q = null } else { l = ae; Q = X } aa.id = R; if (typeof aa.width == D || (!/%$/.test(aa.width) && parseInt(aa.width, 10) < 310)) { aa.width = "310" } if (typeof aa.height == D || (!/%$/.test(aa.height) && parseInt(aa.height, 10) < 137)) { aa.height = "137" } j.title = j.title.slice(0, 47) + " - Flash Player Installation"; var ad = M.ie && M.win ? "ActiveX" : "PlugIn", ac = "MMredirectURL=" + O.location.toString().replace(/&/g, "%26") + "&MMplayerType=" + ad + "&MMdoctitle=" + j.title; if (typeof ab.flashvars != D) { ab.flashvars += "&" + ac } else { ab.flashvars = ac } if (M.ie && M.win && ae.readyState != 4) { var Y = C("div"); X += "SWFObjectNew"; Y.setAttribute("id", X); ae.parentNode.insertBefore(Y, ae); ae.style.display = "none"; (function () { if (ae.readyState == 4) { ae.parentNode.removeChild(ae) } else { setTimeout(arguments.callee, 10) } })() } u(aa, ab, X) } } function p(Y) { if (M.ie && M.win && Y.readyState != 4) { var X = C("div"); Y.parentNode.insertBefore(X, Y); X.parentNode.replaceChild(g(Y), X); Y.style.display = "none"; (function () { if (Y.readyState == 4) { Y.parentNode.removeChild(Y) } else { setTimeout(arguments.callee, 10) } })() } else { Y.parentNode.replaceChild(g(Y), Y) } } function g(ab) { var aa = C("div"); if (M.win && M.ie) { aa.innerHTML = ab.innerHTML } else { var Y = ab.getElementsByTagName(r)[0]; if (Y) { var ad = Y.childNodes; if (ad) { var X = ad.length; for (var Z = 0; Z < X; Z++) { if (!(ad[Z].nodeType == 1 && ad[Z].nodeName == "PARAM") && !(ad[Z].nodeType == 8)) { aa.appendChild(ad[Z].cloneNode(true)) } } } } } return aa } function u(ai, ag, Y) { var X, aa = c(Y); if (M.wk && M.wk < 312) { return X } if (aa) { if (typeof ai.id == D) { ai.id = Y } if (M.ie && M.win) { var ah = ""; for (var ae in ai) { if (ai[ae] != Object.prototype[ae]) { if (ae.toLowerCase() == "data") { ag.movie = ai[ae] } else { if (ae.toLowerCase() == "styleclass") { ah += ' class="' + ai[ae] + '"' } else { if (ae.toLowerCase() != "classid") { ah += " " + ae + '="' + ai[ae] + '"' } } } } } var af = ""; for (var ad in ag) { if (ag[ad] != Object.prototype[ad]) { af += '<param name="' + ad + '" value="' + ag[ad] + '" />' } } aa.outerHTML = '<object classid="clsid:D27CDB6E-AE6D-11cf-96B8-444553540000"' + ah + ">" + af + "</object>"; N[N.length] = ai.id; X = c(ai.id) } else { var Z = C(r); Z.setAttribute("type", q); for (var ac in ai) { if (ai[ac] != Object.prototype[ac]) { if (ac.toLowerCase() == "styleclass") { Z.setAttribute("class", ai[ac]) } else { if (ac.toLowerCase() != "classid") { Z.setAttribute(ac, ai[ac]) } } } } for (var ab in ag) { if (ag[ab] != Object.prototype[ab] && ab.toLowerCase() != "movie") { e(Z, ab, ag[ab]) } } aa.parentNode.replaceChild(Z, aa); X = Z } } return X } function e(Z, X, Y) { var aa = C("param"); aa.setAttribute("name", X); aa.setAttribute("value", Y); Z.appendChild(aa) } function y(Y) { var X = c(Y); if (X && X.nodeName == "OBJECT") { if (M.ie && M.win) { X.style.display = "none"; (function () { if (X.readyState == 4) { b(Y) } else { setTimeout(arguments.callee, 10) } })() } else { X.parentNode.removeChild(X) } } } function b(Z) { var Y = c(Z); if (Y) { for (var X in Y) { if (typeof Y[X] == "function") { Y[X] = null } } Y.parentNode.removeChild(Y) } } function c(Z) { var X = null; try { X = j.getElementById(Z) } catch (Y) { } return X } function C(X) { return j.createElement(X) } function i(Z, X, Y) { Z.attachEvent(X, Y); I[I.length] = [Z, X, Y] } function F(Z) { var Y = M.pv, X = Z.split("."); X[0] = parseInt(X[0], 10); X[1] = parseInt(X[1], 10) || 0; X[2] = parseInt(X[2], 10) || 0; return (Y[0] > X[0] || (Y[0] == X[0] && Y[1] > X[1]) || (Y[0] == X[0] && Y[1] == X[1] && Y[2] >= X[2])) ? true : false } function v(ac, Y, ad, ab) { if (M.ie && M.mac) { return } var aa = j.getElementsByTagName("head")[0]; if (!aa) { return } var X = (ad && typeof ad == "string") ? ad : "screen"; if (ab) { n = null; G = null } if (!n || G != X) { var Z = C("style"); Z.setAttribute("type", "text/css"); Z.setAttribute("media", X); n = aa.appendChild(Z); if (M.ie && M.win && typeof j.styleSheets != D && j.styleSheets.length > 0) { n = j.styleSheets[j.styleSheets.length - 1] } G = X } if (M.ie && M.win) { if (n && typeof n.addRule == r) { n.addRule(ac, Y) } } else { if (n && typeof j.createTextNode != D) { n.appendChild(j.createTextNode(ac + " {" + Y + "}")) } } } function w(Z, X) { if (!m) { return } var Y = X ? "visible" : "hidden"; if (J && c(Z)) { c(Z).style.visibility = Y } else { v("#" + Z, "visibility:" + Y) } } function L(Y) { var Z = /[\\\"<>\.;]/; var X = Z.exec(Y) != null; return X && typeof encodeURIComponent != D ? encodeURIComponent(Y) : Y } var d = function () { if (M.ie && M.win) { window.attachEvent("onunload", function () { var ac = I.length; for (var ab = 0; ab < ac; ab++) { I[ab][0].detachEvent(I[ab][1], I[ab][2]) } var Z = N.length; for (var aa = 0; aa < Z; aa++) { y(N[aa]) } for (var Y in M) { M[Y] = null } M = null; for (var X in swfobject) { swfobject[X] = null } swfobject = null }) } }(); return { registerObject: function (ab, X, aa, Z) { if (M.w3 && ab && X) { var Y = {}; Y.id = ab; Y.swfVersion = X; Y.expressInstall = aa; Y.callbackFn = Z; o[o.length] = Y; w(ab, false) } else { if (Z) { Z({ success: false, id: ab }) } } }, getObjectById: function (X) { if (M.w3) { return z(X) } }, embedSWF: function (ab, ah, ae, ag, Y, aa, Z, ad, af, ac) { var X = { success: false, id: ah }; if (M.w3 && !(M.wk && M.wk < 312) && ab && ah && ae && ag && Y) { w(ah, false); K(function () { ae += ""; ag += ""; var aj = {}; if (af && typeof af === r) { for (var al in af) { aj[al] = af[al] } } aj.data = ab; aj.width = ae; aj.height = ag; var am = {}; if (ad && typeof ad === r) { for (var ak in ad) { am[ak] = ad[ak] } } if (Z && typeof Z === r) { for (var ai in Z) { if (typeof am.flashvars != D) { am.flashvars += "&" + ai + "=" + Z[ai] } else { am.flashvars = ai + "=" + Z[ai] } } } if (F(Y)) { var an = u(aj, am, ah); if (aj.id == ah) { w(ah, true) } X.success = true; X.ref = an } else { if (aa && A()) { aj.data = aa; P(aj, am, ah, ac); return } else { w(ah, true) } } if (ac) { ac(X) } }) } else { if (ac) { ac(X) } } }, switchOffAutoHideShow: function () { m = false }, ua: M, getFlashPlayerVersion: function () { return { major: M.pv[0], minor: M.pv[1], release: M.pv[2] } }, hasFlashPlayerVersion: F, createSWF: function (Z, Y, X) { if (M.w3) { return u(Z, Y, X) } else { return undefined } }, showExpressInstall: function (Z, aa, X, Y) { if (M.w3 && A()) { P(Z, aa, X, Y) } }, removeSWF: function (X) { if (M.w3) { y(X) } }, createCSS: function (aa, Z, Y, X) { if (M.w3) { v(aa, Z, Y, X) } }, addDomLoadEvent: K, addLoadEvent: s, getQueryParamValue: function (aa) { var Z = j.location.search || j.location.hash; if (Z) { if (/\?/.test(Z)) { Z = Z.split("?")[1] } if (aa == null) { return L(Z) } var Y = Z.split("&"); for (var X = 0; X < Y.length; X++) { if (Y[X].substring(0, Y[X].indexOf("=")) == aa) { return L(Y[X].substring((Y[X].indexOf("=") + 1))) } } } return "" }, expressInstallCallback: function () { if (a) { var X = c(R); if (X && l) { X.parentNode.replaceChild(l, X); if (Q) { w(Q, true); if (M.ie && M.win) { l.style.display = "block" } } if (E) { E(B) } } a = false } } } }();
(function () { if (window.WEB_SOCKET_FORCE_FLASH) { } else { if (window.WebSocket) { return } else { if (window.MozWebSocket) { window.WebSocket = MozWebSocket; return } } } var a; if (window.WEB_SOCKET_LOGGER) { a = WEB_SOCKET_LOGGER } else { if (window.console && window.console.log && window.console.error) { a = window.console } else { a = { log: function () { }, error: function () { } } } } if (swfobject.getFlashPlayerVersion().major < 10) { a.error("Flash Player >= 10.0.0 is required."); return } if (location.protocol == "file:") { a.error("WARNING: web-socket-js doesn't work in file:///... URL " + "unless you set Flash Security Settings properly. " + "Open the page via Web server i.e. http://...") } window.WebSocket = function (d, e, c, g, f) { var b = this; b.__id = WebSocket.__nextId++; WebSocket.__instances[b.__id] = b; b.readyState = WebSocket.CONNECTING; b.bufferedAmount = 0; b.__events = {}; if (!e) { e = [] } else { if (typeof e == "string") { e = [e] } } b.__createTask = setTimeout(function () { WebSocket.__addTask(function () { b.__createTask = null; WebSocket.__flash.create(b.__id, d, e, c || null, g || 0, f || null) }) }, 0) }; WebSocket.prototype.send = function (c) { if (this.readyState == WebSocket.CONNECTING) { throw "INVALID_STATE_ERR: Web Socket connection has not been established" } var b = WebSocket.__flash.send(this.__id, encodeURIComponent(c)); if (b < 0) { return true } else { this.bufferedAmount += b; return false } }; WebSocket.prototype.close = function () { if (this.__createTask) { clearTimeout(this.__createTask); this.__createTask = null; this.readyState = WebSocket.CLOSED; return } if (this.readyState == WebSocket.CLOSED || this.readyState == WebSocket.CLOSING) { return } this.readyState = WebSocket.CLOSING; WebSocket.__flash.close(this.__id) }; WebSocket.prototype.addEventListener = function (c, d, b) { if (!(c in this.__events)) { this.__events[c] = [] } this.__events[c].push(d) }; WebSocket.prototype.removeEventListener = function (e, f, b) { if (!(e in this.__events)) { return } var d = this.__events[e]; for (var c = d.length - 1; c >= 0; --c) { if (d[c] === f) { d.splice(c, 1); break } } }; WebSocket.prototype.dispatchEvent = function (e) { var c = this.__events[e.type] || []; for (var b = 0; b < c.length; ++b) { c[b](e) } var d = this["on" + e.type]; if (d) { d.apply(this, [e]) } }; WebSocket.prototype.__handleEvent = function (d) { if ("readyState" in d) { this.readyState = d.readyState } if ("protocol" in d) { this.protocol = d.protocol } var b; if (d.type == "open" || d.type == "error") { b = this.__createSimpleEvent(d.type) } else { if (d.type == "close") { b = this.__createSimpleEvent("close"); b.wasClean = d.wasClean ? true : false; b.code = d.code; b.reason = d.reason } else { if (d.type == "message") { var c = decodeURIComponent(d.message); b = this.__createMessageEvent("message", c) } else { throw "unknown event type: " + d.type } } } this.dispatchEvent(b) }; WebSocket.prototype.__createSimpleEvent = function (b) { if (document.createEvent && window.Event) { var c = document.createEvent("Event"); c.initEvent(b, false, false); return c } else { return { type: b, bubbles: false, cancelable: false } } }; WebSocket.prototype.__createMessageEvent = function (b, d) { if (window.MessageEvent && typeof (MessageEvent) == "function" && !window.opera) { return new MessageEvent("message", { "view": window, "bubbles": false, "cancelable": false, "data": d }) } else { if (document.createEvent && window.MessageEvent && !window.opera) { var c = document.createEvent("MessageEvent"); c.initMessageEvent("message", false, false, d, null, null, window, null); return c } else { return { type: b, data: d, bubbles: false, cancelable: false } } } }; WebSocket.CONNECTING = 0; WebSocket.OPEN = 1; WebSocket.CLOSING = 2; WebSocket.CLOSED = 3; WebSocket.__isFlashImplementation = true; WebSocket.__initialized = false; WebSocket.__flash = null; WebSocket.__instances = {}; WebSocket.__tasks = []; WebSocket.__nextId = 0; WebSocket.loadFlashPolicyFile = function (b) { WebSocket.__addTask(function () { WebSocket.__flash.loadManualPolicyFile(b) }) }; WebSocket.__initialize = function () { if (WebSocket.__initialized) { return } WebSocket.__initialized = true; if (WebSocket.__swfLocation) { window.WEB_SOCKET_SWF_LOCATION = WebSocket.__swfLocation } if (!window.WEB_SOCKET_SWF_LOCATION) { a.error("[WebSocket] set WEB_SOCKET_SWF_LOCATION to location of WebSocketMain.swf"); return } if (!window.WEB_SOCKET_SUPPRESS_CROSS_DOMAIN_SWF_ERROR && !WEB_SOCKET_SWF_LOCATION.match(/(^|\/)WebSocketMainInsecure\.swf(\?.*)?$/) && WEB_SOCKET_SWF_LOCATION.match(/^\w+:\/\/([^\/]+)/)) { var d = RegExp.$1; if (location.host != d) { a.error("[WebSocket] You must host HTML and WebSocketMain.swf in the same host " + "('" + location.host + "' != '" + d + "'). " + "See also 'How to host HTML file and SWF file in different domains' section " + "in README.md. If you use WebSocketMainInsecure.swf, you can suppress this message " + "by WEB_SOCKET_SUPPRESS_CROSS_DOMAIN_SWF_ERROR = true;") } } var b = document.createElement("div"); b.id = "webSocketContainer"; b.style.position = "absolute"; if (WebSocket.__isFlashLite()) { b.style.left = "0px"; b.style.top = "0px" } else { b.style.left = "-100px"; b.style.top = "-100px" } var c = document.createElement("div"); c.id = "webSocketFlash"; b.appendChild(c); document.body.appendChild(b); swfobject.embedSWF(WEB_SOCKET_SWF_LOCATION, "webSocketFlash", "1", "1", "10.0.0", null, null, { hasPriority: true, swliveconnect: true, allowScriptAccess: "always" }, null, function (f) { if (!f.success) { a.error("[WebSocket] swfobject.embedSWF failed") } }) }; WebSocket.__onFlashInitialized = function () { setTimeout(function () { WebSocket.__flash = document.getElementById("webSocketFlash"); WebSocket.__flash.setCallerUrl(location.href); WebSocket.__flash.setDebug(!!window.WEB_SOCKET_DEBUG); for (var b = 0; b < WebSocket.__tasks.length; ++b) { WebSocket.__tasks[b]() } WebSocket.__tasks = [] }, 0) }; WebSocket.__onFlashEvent = function () { setTimeout(function () { try { var c = WebSocket.__flash.receiveEvents(); for (var b = 0; b < c.length; ++b) { WebSocket.__instances[c[b].webSocketId].__handleEvent(c[b]) } } catch (d) { a.error(d) } }, 0); return true }; WebSocket.__log = function (b) { a.log(decodeURIComponent(b)) }; WebSocket.__error = function (b) { a.error(decodeURIComponent(b)) }; WebSocket.__addTask = function (b) { if (WebSocket.__flash) { b() } else { WebSocket.__tasks.push(b) } }; WebSocket.__isFlashLite = function () { if (!window.navigator || !window.navigator.mimeTypes) { return false } var b = window.navigator.mimeTypes["application/x-shockwave-flash"]; if (!b || !b.enabledPlugin || !b.enabledPlugin.filename) { return false } return b.enabledPlugin.filename.match(/flashlite/i) ? true : false }; if (!window.WEB_SOCKET_DISABLE_AUTO_INITIALIZATION) { swfobject.addDomLoadEvent(function () { WebSocket.__initialize() }) } })();

var ws;
var ws_url;
var isopenEmoji = false;
var tick_heartpac = null;

WEB_SOCKET_SWF_LOCATION = "PLUG/WebSocketMain.swf";
var emoji = {
    "[最右]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/c8/lxhzuiyou_thumb.gif",
    "[泪流满面]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/64/lxhtongku_thumb.gif",
    "[江南style]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/67/gangnamstyle_thumb.gif",
    "[偷乐]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/fa/lxhtouxiao_thumb.gif",
    "[加油啊]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/03/lxhjiayou_thumb.gif",
    "[doge]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/b6/doge_thumb.gif",
    "[喵喵]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/4a/mm_thumb.gif",
    "[笑cry]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/34/xiaoku_thumb.gif",
    "[xkl转圈]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/f4/xklzhuanquan_thumb.gif",
    "[微笑]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/5c/huanglianwx_thumb.gif",
    "[嘻嘻]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/0b/tootha_thumb.gif",
    "[哈哈]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/6a/laugh.gif",
    "[可爱]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/14/tza_thumb.gif",
    "[可怜]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/af/kl_thumb.gif",
    "[挖鼻]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/0b/wabi_thumb.gif",
    "[吃惊]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/f4/cj_thumb.gif",
    "[害羞]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/6e/shamea_thumb.gif",
    "[挤眼]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/c3/zy_thumb.gif",
    "[闭嘴]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/29/bz_thumb.gif",
    "[鄙视]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/71/bs2_thumb.gif",
    "[爱你]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/6d/lovea_thumb.gif",
    "[泪]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/9d/sada_thumb.gif",
    "[偷笑]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/19/heia_thumb.gif",
    "[亲亲]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/8f/qq_thumb.gif",
    "[生病]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/b6/sb_thumb.gif",
    "[太开心]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/58/mb_thumb.gif",
    "[白眼]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/d9/landeln_thumb.gif",
    "[右哼哼]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/98/yhh_thumb.gif",
    "[左哼哼]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/6d/zhh_thumb.gif",
    "[嘘]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/a6/x_thumb.gif",
    "[衰]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/af/cry.gif",
    "[委屈]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/73/wq_thumb.gif",
    "[吐]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/9e/t_thumb.gif",
    "[哈欠]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/cc/haqianv2_thumb.gif",
    "[抱抱]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/27/bba_thumb.gif",
    "[怒]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/7c/angrya_thumb.gif",
    "[疑问]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/5c/yw_thumb.gif",
    "[馋嘴]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/a5/cza_thumb.gif",
    "[拜拜]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/70/88_thumb.gif",
    "[思考]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/e9/sk_thumb.gif",
    "[汗]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/24/sweata_thumb.gif",
    "[困]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/40/kunv2_thumb.gif",
    "[睡]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/96/huangliansj_thumb.gif",
    "[钱]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/90/money_thumb.gif",
    "[失望]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/0c/sw_thumb.gif",
    "[酷]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/40/cool_thumb.gif",
    "[色]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/20/huanglianse_thumb.gif",
    "[哼]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/49/hatea_thumb.gif",
    "[鼓掌]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/36/gza_thumb.gif",
    "[晕]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/d9/dizzya_thumb.gif",
    "[悲伤]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/1a/bs_thumb.gif",
    "[泪]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/9d/sada_thumb.gif",
    "[偷笑]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/19/heia_thumb.gif",
    "[抓狂]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/62/crazya_thumb.gif",
    "[黑线]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/91/h_thumb.gif",
    "[阴险]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/6d/yx_thumb.gif",
    "[怒骂]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/60/numav2_thumb.gif",
    "[互粉]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/89/hufen_thumb.gif",
    "[心]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/40/hearta_thumb.gif",
    "[伤心]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/ea/unheart.gif",
    "[猪头]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/58/pig.gif",
    "[熊猫]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/6e/panda_thumb.gif",
    "[兔子]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/81/rabbit_thumb.gif",
    "[ok]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/d6/ok_thumb.gif",
    "[耶]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/d9/ye_thumb.gif",
    "[good]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/d8/good_thumb.gif",
    "[no]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/ae/buyao_org.gif",
    "[赞]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/d0/z2_thumb.gif",
    "[来]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/40/come_thumb.gif",
    "[弱]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/d8/sad_thumb.gif",
    "[草泥马]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/7a/shenshou_thumb.gif",
    "[神马]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/60/horse2_thumb.gif",
    "[囧]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/15/j_thumb.gif",
    "[浮云]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/bc/fuyun_thumb.gif",
    "[给力]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/1e/geiliv2_thumb.gif",
    "[围观]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/f2/wg_thumb.gif",
    "[威武]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/70/vw_thumb.gif",
    "[奥特曼]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/bc/otm_thumb.gif",
    "[礼物]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/c4/liwu_thumb.gif",
    "[钟]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/d3/clock_thumb.gif",
    "[话筒]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/9f/huatongv2_thumb.gif",
    "[蜡烛]": "http://img.t.sinajs.cn/t4/appstyle/expression/ext/normal/d9/lazhuv2_thumb.gif"
};

//jquery扩展
$.fn.extend({
    insertAtCaret: function (myValue) {
        var $t = $(this)[0];
        if (document.selection) {
            this.focus();
            sel = document.selection.createRange();
            sel.text = myValue;
            this.focus();
        } else
            if ($t.selectionStart || $t.selectionStart == '0') {
                var startPos = $t.selectionStart;
                var endPos = $t.selectionEnd;
                var scrollTop = $t.scrollTop;
                $t.value = $t.value.substring(0, startPos) + myValue + $t.value.substring(endPos, $t.value.length);
                this.focus();
                $t.selectionStart = startPos + myValue.length;
                $t.selectionEnd = startPos + myValue.length;
                $t.scrollTop = scrollTop;
            } else {
                this.value += myValue;
                this.focus();
            }
    }//在光标处插入文本
});

$(function () {
    fun_initEmoji();
    $("#div_msg").width(window.innerWidth - 530);
    $("#div_msg").height(window.innerHeight - 200);
    $("#inp_send").keydown(function (e) {
        if (e.keyCode == 13 && e.ctrlKey) {
            $("#btn_send").trigger("click");
            return false;
        }
    });
});

//置换js语句
function NoHtml(t) {
    t = t.replace(/({|})/g, '');
    t = t.replace(/</g, '&lt;');
    t = t.replace(/>/g, '&gt;');
    return t;
}

//启动或停止30秒一次心跳包
function Heart(start) {
    if (start) {
        tick_heartpac = setInterval(function () {
            ws.send("@heart");
        }, 30000);
    }
    else {
        clearInterval(tick_heartpac);
        tick_heartpac = null;
    }
}

//定义表情包
function fun_initEmoji() {
    var emoji_html = "";
    for (var i in emoji) {
        emoji_html += "<li class='list_emoji' id='emoji_" + i + "' onclick='click_emojiItems(id);'  title='" + i + "'><img src='" + emoji[i] + "' /></li>";
    }
    $("#emoji").popover({ html: "true", content: emoji_html });
    $("#emoji").click(function () {
        if (!isopenEmoji) {
            isopenEmoji = true;
            $("#emoji").popover('show');
        }
        else {
            isopenEmoji = false;
            $("#emoji").popover('hide');
        }
    });
}

//表情项的点击
function click_emojiItems(id) {
    $("#inp_send").insertAtCaret(id.split('_')[1]);
}

//定义socket
function fun_initWebSocket() {
    ws_url = $.trim($("#inp_url").val()).toLocaleLowerCase();
    if (ws_url) {
        $("#btn_conn").attr('disabled', true);
        $("#btn_close").attr('disabled', false);
        try {
            ws = new WebSocket($.trim($("#inp_url").val()));
            output("等待服务器握手包...", 1);
            ws.onopen = function (e) {
                output("收到服务器握手包.", 1);
                output("连接已建立，正在等待数据...", 0);
                Heart(true);
            };
            ws.onmessage = function (e) {
                output(chg_emoji(e.data), 0);
            };
            ws.onclose = function (e) {
                $("#btn_conn").attr('disabled', false);
                $("#btn_close").attr('disabled', true);
                Heart(false);
                output("和服务器断开连接！", 0);
            };
        }
        catch (e) {
            $("#btn_conn").attr('disabled', false);
            $("#btn_close").attr('disabled', true);
            output("ws的地址错误,请重新输入！", 1);
        }
    }
}

//关闭连接
function fun_close() {
    ws.close();
    Heart(false);
}

//向中心发送数据
function fun_sendto() {
    if (isopenEmoji)
        $("#emoji").trigger("click");
    var msg = $.trim($("#inp_send").val());
    if (msg == "")
        return;
    if (ws && ws.readyState == 1) {
        Heart(false);
        ws.send(msg);
        output(NoHtml(chg_emoji(msg)), 1);
        $("#inp_send").val("");
        Heart(true);
    }
    else
        alert('连接已经断开！');
}

//转换文字为表情
function chg_emoji(str) {
    var em = str.match(/\[.*?]/gi);
    if (em) {
        for (var i = 0; i < em.length; i++) {
            str = str.replace(em[i], "<img src='" + emoji[em[i]] + "' />");
        }
    }
    return str;
}

//显示信息
function output(str, sign) {
    var time = new Date();
    var color = "blue";
    var user = "服务器";
    if (sign == 1) {
        color = "green";
        user = "你";
    }
    var str_time = "<div style='color:" + color + "'>" + user + " " + time.getHours() + ":" + time.getMinutes() + ":" + time.getSeconds() + "</div>";
    $("#div_msg").append("<div style='margin-bottom:10px;position:relative;left:0px;'>" + str_time + str + "</div>");
    $('#div_msg').scrollTop($('#div_msg')[0].scrollHeight);
}