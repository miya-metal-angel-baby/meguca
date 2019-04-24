// Selects and loads the client files and polyfills, if any. Use only ES5.

(function () {
	// Check if the client is an automated crawler
	var isBot,
		botStrings = [
			"bot", "googlebot", "crawler", "spider", "robot", "crawling"
		];

	for (var i = 0; i < botStrings.length; i++) {
		if (navigator.userAgent.indexOf(botStrings[i]) !== -1) {
			isBot = true;
			break;
		}
	}

	// Display mature content warning
/*
	if (!isBot && !localStorage.getItem("termsAccepted")) {
		var confirmText =
			"To access this website you understand and agree to the following:\n\n" +
			"1. The content of this website is for mature audiences only and may not be suitable for minors. If you are a minor or it is illegal for you to access mature images and language, do not proceed.\n\n" +
			"2. This website is presented to you AS IS, with no warranty, express or implied. By proceeding you agree not to hold the owner(s) of the website responsible for any damages from your use of the website, and you understand that the content posted is not owned or generated by the website, but rather by the website's users.";

		if (!confirm(confirmText)) {
			location.href = "http://www.gaiaonline.com/";
			return;
		}

		localStorage.setItem("termsAccepted", "true");
	}
*/

	// Really old browser. Run in noscript mode.
	if (!window.WebAssembly) {
		var ns = document.getElementsByTagName("noscript");

		while (ns.length) { // Collection is live and changes with DOM updates
			var el = ns[0],
				cont = document.createElement("div");
			cont.innerHTML = el.innerHTML;
			el.parentNode.replaceChild(cont, el);
		}

		var bc = document.getElementById("banner-center");
		bc.classList.add("admin");
		bc.innerHTML = "UPDATE YOUR FUCKING BROWSER";
		return;
	}

	var scriptCount = 0,
		polyfills = [];

	var DOMUpToDate = true,
		DOMMethods = [
			// DOM level 4 methods
			'Element.prototype.remove',
			'Element.prototype.contains',
			'Element.prototype.matches',
			'Element.prototype.after',
			'Element.prototype.before',
			'Element.prototype.append',
			'Element.prototype.prepend',
			'Element.prototype.replaceWith',

			// DOM level 3 query methods
			'Element.prototype.querySelector',
			'Element.prototype.querySelectorAll'
		];

	for (i = 0; i < DOMMethods.length; i++) {
		if (!checkFunction(DOMMethods[i])) {
			DOMUpToDate = false;
			break;
		}
	}

	// Check event listener option support
	if (DOMUpToDate) {
		var s = "var a = document.createElement(\"a\");" +
			"var ctr = 0;" +
			"a.addEventListener(\"click\", () => ctr++, {once: true});" +
			"a.click(); a.click();" +
			"return ctr === 1;";
		DOMUpToDate = check(s);
	}

	if (!DOMUpToDate) {
		polyfills.push('js/vendor/dom4');
	}

	// Stdlib functions and methods
	var stdlibTests = [
		"Set",
		"Map",
		'Promise',
		"Symbol",
		"Array.from",
		'Array.prototype.includes',
		"String.prototype.includes",
		"String.prototype.repeat"
	];

	for (i = 0; i < stdlibTests.length; i++) {
		if (!checkFunction(stdlibTests[i])) {
			polyfills.push("js/vendor/core.min");
			break;
		}
	}

	// Remove prefixes on Web Crypto API for Safari
	if (!checkFunction("window.crypto.subtle.digest")) {
		window.crypto.subtle = window.crypto.webkitSubtle;
	}

	var wasm = /[\?&]wasm=true/.test(location.search),
		head = document.getElementsByTagName('head')[0];

	if (polyfills.length) {
		for (i = 0; i < polyfills.length; i++) {
			scriptCount++;
			loadScript(polyfills[i]).onload = checkAllLoaded;
		}
	} else {
		loadClient();
	}

	// Check for browser compatibility by trying to detect some ES6 features
	function check(func) {
		try {
			// Using `eval` can expose this code to injection attacks
			// and creates a new instance of the javascript interpreter (W061)
			return eval('(function(){' + func + '})()');
		} catch (e) {
			return false;
		}
	}

	// Check if a browser API function is defined
	function checkFunction(func) {
		try {
			// See comment on line 134
			return typeof eval(func) === 'function';
		} catch (e) {
			return false;
		}
	}

	function checkAllLoaded() {
		// This function might be called multiple times. Only load the client,
		// when all polyfills are loaded.
		if (--scriptCount === 0) {
			loadClient();
		}
	}

	function loadScript(path) {
		var script = document.createElement('script');
		script.type = 'text/javascript';
		script.src = '/assets/' + path + '.js';
		head.appendChild(script);
		return script;
	}

	function loadClient() {
		// Iterable NodeList
		if (!checkFunction('NodeList.prototype[Symbol.iterator]')) {
			NodeList.prototype[Symbol.iterator] =
				Array.prototype[Symbol.iterator];
		}

		if (wasm) {
			window.Module = {};
			fetch("/assets/wasm/main.wasm").then(function (res) {
				return res.arrayBuffer();
			}).then(function (bytes) {
				// TODO: Parallel downloads of main.js and main.wasm
				var script = document.createElement('script');
				script.src = "/assets/wasm/main.js";
				Module.wasmBinary = bytes;
				document.head.appendChild(script);
			});
		} else {
			loadScript("js/main").onload = function () {
				require("main");
			};
		}

		if ('serviceWorker' in navigator && (
			location.protocol === "https:" ||
			location.hostname === "localhost"
		)) {
			navigator.serviceWorker
				.register("/assets/js/scripts/worker.js", { scope: "/" })
				.catch(function (err) {
					throw err;
				});
		}
	}
})();
