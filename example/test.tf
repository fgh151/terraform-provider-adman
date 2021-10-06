terraform {
  required_providers {
    adman = {
      source = "openitstudio.ru/dns/adman"
      version = "0.0.1"
    }
  }
}

provider "adman" {
  login = "login"
  mdpass = "mdpass"
}

resource "adman_dns_zone" "test_com" {
  domain  = "test.com"
}

resource "adman_dns_zone_record" "a_a" {
  zone = adman_dns_zone.test_com.domain
  host            = "A"
  type            = "TXT"
  value           = "11.22.33.44"
  ttl             = 10
  external_id     = ""
  additional_info = ""
}
