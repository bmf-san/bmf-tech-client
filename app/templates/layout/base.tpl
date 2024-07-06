{{ define "base" }}
<!DOCTYPE html>
<html lang="ja">
<head>
    {{ template "meta" .Meta }}
    <link rel="icon" href="/favicon.ico" />
    <link rel="stylesheet" href="https://unpkg.com/sea.css/dist/sea.min.css">
    <link rel="stylesheet" href="/css/style.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/gh/highlightjs/cdn-release@11.5.1/build/styles/monokai.min.css">
    <script src="https://cdnjs.cloudflare.com/ajax/libs/highlight.js/11.5.1/highlight.min.js"></script>
    <script>hljs.initHighlightingOnLoad();</script>
    <!-- Global site tag (gtag.js) - Google Analytics -->
    <script async src="https://www.googletagmanager.com/gtag/js?id=G-784B55NW88"></script>
    <script async>
        window.dataLayer = window.dataLayer || [];
        function gtag() { dataLayer.push(arguments); }
        gtag('js', new Date());

        gtag('config', 'G-784B55NW88');
    </script>
    <!-- InMobi Choice. Consent Manager Tag v3.0 (for TCF 2.2) -->
    <script type="text/javascript" async=true>
    (function() {
      var host = "www.themoneytizer.com";
      var element = document.createElement('script');
      var firstScript = document.getElementsByTagName('script')[0];
      var url = 'https://cmp.inmobi.com'
        .concat('/choice/', '6Fv0cGNfc_bw8', '/', host, '/choice.js?tag_version=V3');
      var uspTries = 0;
      var uspTriesLimit = 3;
      element.async = true;
      element.type = 'text/javascript';
      element.src = url;

      firstScript.parentNode.insertBefore(element, firstScript);

      function makeStub() {
        var TCF_LOCATOR_NAME = '__tcfapiLocator';
        var queue = [];
        var win = window;
        var cmpFrame;

        function addFrame() {
          var doc = win.document;
          var otherCMP = !!(win.frames[TCF_LOCATOR_NAME]);

          if (!otherCMP) {
            if (doc.body) {
              var iframe = doc.createElement('iframe');

              iframe.style.cssText = 'display:none';
              iframe.name = TCF_LOCATOR_NAME;
              doc.body.appendChild(iframe);
            } else {
              setTimeout(addFrame, 5);
            }
          }
          return !otherCMP;
        }

        function tcfAPIHandler() {
          var gdprApplies;
          var args = arguments;

          if (!args.length) {
            return queue;
          } else if (args[0] === 'setGdprApplies') {
            if (
              args.length > 3 &&
              args[2] === 2 &&
              typeof args[3] === 'boolean'
            ) {
              gdprApplies = args[3];
              if (typeof args[2] === 'function') {
                args[2]('set', true);
              }
            }
          } else if (args[0] === 'ping') {
            var retr = {
              gdprApplies: gdprApplies,
              cmpLoaded: false,
              cmpStatus: 'stub'
            };

            if (typeof args[2] === 'function') {
              args[2](retr);
            }
          } else {
            if(args[0] === 'init' && typeof args[3] === 'object') {
              args[3] = Object.assign(args[3], { tag_version: 'V3' });
            }
            queue.push(args);
          }
        }

        function postMessageEventHandler(event) {
          var msgIsString = typeof event.data === 'string';
          var json = {};

          try {
            if (msgIsString) {
              json = JSON.parse(event.data);
            } else {
              json = event.data;
            }
          } catch (ignore) {}

          var payload = json.__tcfapiCall;

          if (payload) {
            window.__tcfapi(
              payload.command,
              payload.version,
              function(retValue, success) {
                var returnMsg = {
                  __tcfapiReturn: {
                    returnValue: retValue,
                    success: success,
                    callId: payload.callId
                  }
                };
                if (msgIsString) {
                  returnMsg = JSON.stringify(returnMsg);
                }
                if (event && event.source && event.source.postMessage) {
                  event.source.postMessage(returnMsg, '*');
                }
              },
              payload.parameter
            );
          }
        }

        while (win) {
          try {
            if (win.frames[TCF_LOCATOR_NAME]) {
              cmpFrame = win;
              break;
            }
          } catch (ignore) {}

          if (win === window.top) {
            break;
          }
          win = win.parent;
        }
        if (!cmpFrame) {
          addFrame();
          win.__tcfapi = tcfAPIHandler;
          win.addEventListener('message', postMessageEventHandler, false);
        }
      };

      makeStub();

      var uspStubFunction = function() {
        var arg = arguments;
        if (typeof window.__uspapi !== uspStubFunction) {
          setTimeout(function() {
            if (typeof window.__uspapi !== 'undefined') {
              window.__uspapi.apply(window.__uspapi, arg);
            }
          }, 500);
        }
      };

      var checkIfUspIsReady = function() {
        uspTries++;
        if (window.__uspapi === uspStubFunction && uspTries < uspTriesLimit) {
          console.warn('USP is not accessible');
        } else {
          clearInterval(uspInterval);
        }
      };

      if (typeof window.__uspapi === 'undefined') {
        window.__uspapi = uspStubFunction;
        var uspInterval = setInterval(checkIfUspIsReady, 6000);
      }
    })();
    </script>
    <!-- End InMobi Choice. Consent Manager Tag v3.0 (for TCF 2.2) -->
</head>
<body>
    <header>
        <div>
            <nav class="nav sp-hidden">
                <div class="col-nav">
                    <div class="nav-left">
                        <a class="nav-link-logo color-text" href="/"><b>bmf-tech</b></a>
                    </div>
                    <div class="nav-right">
                        <a class="nav-link color-text" href="/posts">記事</a>
                        <a class="nav-link color-text" href="/categories">カテゴリ</a>
                        <a class="nav-link color-text" href="/tags">タグ</a>
                        <a class="nav-link color-text" href="/profile">プロフィール</a>
                    </div>
                </div>
            </nav>
            <nav class="nav pc-hidden">
                <div class="col-nav">
                    <div class="nav-left">
                        <a class="nav-link text-decoration-none color-text" href="/">bmf-tech</a>
                    </div>
                    <div class="nav-right">
                        <div id="nav-sp-drawer">
                            <input id="nav-sp-input" type="checkbox" class="sp-hidden">
                            <label id="nav-sp-open" for="nav-sp-input"><span></span></label>
                            <label class="sp-hidden" id="nav-sp-close" for="nav-sp-input"><span></span></label>
                            <div id="nav-sp-content">
                                <h1>bmf-tech</h1>
                                <a class="nav-link color-text" href="/posts">記事</a>
                                <a class="nav-link color-text" href="/categories">カテゴリ</a>
                                <a class="nav-link color-text" href="/tags">タグ</a>
                                <a class="nav-link color-text" href="/profile">プロフィール</a>
                            </div>
                        </div>
                    </div>
                </div>
            </nav>
        </div>
    </header>
    {{ template "headline" . }}
    {{ template "body" . }}
    <footer class="sticky-footer margin-top-2rem">
        <div class="container">
            <div class="row text-align-center font-size-small">
                <div class="col">
                    <a class="color-text-reverse" href="https://twitter.com/bmf_san">Twitter</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="https://github.com/bmf-san">Github</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="/sitemap">サイトマップ</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="/feed">フィード</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="/support">サポート</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="https://forms.gle/zZamAmNEtHZGnwZY6">お問い合わせ</a>
                </div>
                <div class="col">
                    <a class="color-text-reverse" href="/privacy_policy">プライバシーポリシー</a>
                </div>
            </div>
            <hr>
            <div class="row text-align-center">
                <div class="col">
                    <small>bmf-tech.comではサイトの運営費やbmf-sanの技術研鑽のために広告の運用を行っています。</small>
                    <br>
                    <small>©{{ year }} Kenta Takeuchi</small>
                </div>
            </div>
        </div>
    </footer>
    {{ block "script" . }}{{ end }}
</body>
</html>
{{ end }}


