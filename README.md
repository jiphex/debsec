# Debsec

Tool for searching the debian security tracker for a bunch of CVEs in bulk

## Usage
    May% echo CVE-2014-3566 > ./cves
    May% ./debsec -distro wheezy-security < cves
    2015/03/20 17:20:06 WARNING: CVE-2014-3566 is open on wheezy-security for pkg gnutls26: https://security-tracker.debian.org/tracker/redirect/CVE-2014-3566
    2015/03/20 17:20:06 WARNING: CVE-2014-3566 is open on wheezy-security for pkg chromium-browser: https://security-tracker.debian.org/tracker/redirect/CVE-2014-3566
    2015/03/20 17:20:06 WARNING: CVE-2014-3566 is open on wheezy-security for pkg openssl: https://security-tracker.debian.org/tracker/redirect/CVE-2014-3566
    2015/03/20 17:20:06 WARNING: CVE-2014-3566 is open on wheezy-security for pkg nss: https://security-tracker.debian.org/tracker/redirect/CVE-2014-3566
    2015/03/20 17:20:06 WARNING: CVE-2014-3566 is open on wheezy-security for pkg polarssl: https://security-tracker.debian.org/tracker/redirect/CVE-2014-3566
