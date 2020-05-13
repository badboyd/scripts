
import base64
import datetime
import hashlib
import hmac
import time

import os
import sys
import json
import time

import base64
import datetime
import urllib
import requests


GCS_API_ENDPOINT = 'https://storage.googleapis.com'
hmac_key = 'GOOG1EZZDTPVEPPYI4CEXAXHMJ3YBP3Q23LQ2N3CI2HC5AJFYJYH6RLAG2BMQ'
hmac_secret = '0CVQiJq4Om/47hqADKVZK08E0ggmXfHd2lq8x51z'
BUCKET_NAME = 'chotot-photo-staging'
OBJECT_NAME = 'd22c8bc8caba71b28fce4200703cdec2-2633799274031022595.jpg'

# expiration = datetime.datetime.now() +  datetime.timedelta(seconds=600)
expiration = int(1570710240)

def _Base64Sign(url_to_sign):
        digest = hmac.new(
          hmac_secret, url_to_sign.encode('utf-8'), hashlib.sha1).digest()

        signature = base64.standard_b64encode(digest).decode('utf-8')
        print 'signature: ' + signature
        return signature


def _MakeSignatureString(verb, path, content_md5, content_type):
      signature_string = ('{verb}\n'
                          '{content_md5}\n'
                          '{content_type}\n'
                          '{expiration}\n'
                          '{resource}')
      return signature_string.format(verb=verb,
                                     content_md5=content_md5,
                                     content_type=content_type,
                                     expiration=expiration,
                                     resource=path)

def MakeUrl(verb, path, content_type='', content_md5=''):
      signature_string = _MakeSignatureString(verb, path, content_md5,
                                                   content_type)
      print signature_string
      signature_signed = urllib.quote(_Base64Sign(signature_string))
      print signature_signed

      signed_url = "https://cdn.badboyd.com/" + \
              BUCKET_NAME + "/" + OBJECT_NAME + "?GoogleAccessId=" + \
              hmac_key + "&Expires=" + str(expiration) + \
              "&Signature=" + signature_signed

      return signed_url


file_path = '/%s/%s' % (BUCKET_NAME, OBJECT_NAME)

print "GET"
u =  MakeUrl("GET",file_path)
print u
