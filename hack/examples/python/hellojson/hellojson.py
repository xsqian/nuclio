# xsqian

import json

def handler(context, event):
    # slip the try and catch
    req = event.body.decode('utf-8').strip()
    # business logics here to change the shape of the data, skip for now
    res = json.loads(req)
    context.logger.info(f'This is the response')

    return context.Response(body=res,
                            headers={},
                            content_type='text/plain',
                            status_code=200)