import base64
import hashlib
import hmac
import textwrap

class Iris(object):

    def __init__(self, base_url="https://cdn.chotot.org",
        key="xxxx", salt="xxxx"):
        self._key = key.decode('hex')
        self._salt = salt.decode('hex')
        self._base_url = base_url

    def gen_object_id(self, id):
        id_without_ext = id.replace('.jpg', '')
        if len(id_without_ext) <= 10:
            return id
        return "{id_hash}-{id}".format(
            id_hash=hashlib.md5(id_without_ext.encode()).hexdigest(),
            id=id
        )

    def gen_url(self, id, preset):
        path = "/preset:{preset}/plain/{id}".format(
            preset=preset,
            id=self.gen_object_id(id)
        ).encode()

        digest = hmac.new(self._key, msg=self._salt+path, digestmod=hashlib.sha256).digest()
        protection = base64.urlsafe_b64encode(digest).rstrip(b"=")

        final_path = b'%s%s' % (protection, path)

        return '%s/%s' % (
            self._base_url,
            final_path.decode(),
        )
