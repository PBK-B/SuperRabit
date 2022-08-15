package utils

import "strings"

type PlistUtil struct{}

var template = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>items</key>
    <array>
        <dict>
            <key>assets</key>
            <array>
                <dict>
                    <key>kind</key>
                    <string>software-package</string>
                    <key>url</key>
                    <string>@@ipaurl@@</string>
                </dict>
               
            </array>
            <key>metadata</key>
            <dict>
                <key>bundle-identifier</key>
                <string>@@bundleid@@</string>
                <key>bundle-version</key>
                <string>@@version@@</string>
                <key>kind</key>
                <string>software</string>
                <key>subtitle</key>
                <string>App Subtitle</string>
                <key>title</key>
                <string>@@title@@</string>
            </dict>
        </dict>
    </array>
</dict>
</plist>`

func (p PlistUtil) GenPlistContent(ipa, bundleid, version, name string) string {
	plistStr := strings.Replace(template, "@@ipaurl@@", ipa, -1)
	plistStr = strings.Replace(plistStr, "@@bundleid@@", bundleid, -1)
	plistStr = strings.Replace(plistStr, "@@version@@", version, -1)
	plistStr = strings.Replace(plistStr, "@@title@@", name, -1)
	return plistStr
}

//"itms-services:///?action=download-manifest&url=" + app.PlistUrl
