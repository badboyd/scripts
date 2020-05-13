from iris import Iris
import unirest

def prefetch_images_callback(response):
    print('Fetch image response code: %s' % response.code)
    return

def main():
    iris = Iris("https://cdn.chotot.com", "73757368690a", "626164626f79640a")
    url = iris.gen_url("2664732968888824287.jpg", "raw")
    print(url)
    unirest.get(url, callback=prefetch_images_callback)

if __name__ == '__main__':
    main()
