package main

var webPage = `
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <title>publisher</title>
    </style>
</head>
<body>
    <h1>これはメディアの配信面です</h1>
    <div id="ad">
    </div>
    <script>
        function Callback(ad) {
            if (ad != null && ad != "" && ad != undefined) {
                var target = document.getElementById("ad");
                if (ad.ads != null && ad.ads != "" && ad.ads != undefined) {
                    target.innerHTML = ad.ads;
                }
            }
        }

        // ADNW URL
        var targetBaseUrl = "http:///delivery"

        // Ad request
        var adReq = document.createElement("script");
        adReq.src = targetBaseUrl
        document.getElementById("ad").insertBefore(adReq, null);
    </script>
</body>
</html>
`
