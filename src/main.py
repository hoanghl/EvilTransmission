from fastapi import FastAPI, File, HTTPException, UploadFile

from . import handlers

app = FastAPI()


@app.post("/frame/")
async def root(file: UploadFile = File(...)):
    # return {
    #     "name": file.filename,
    #     "type": file.content_type
    # }

    # logger.debug(f"content_type: {file.content_type}")

    if file.content_type is None:
        raise HTTPException(status_code=400, detail="Incorrect data type")

    if file.content_type in ["image/png", "image/jpg", "image/jpeg"]:
        out = handlers.handle_image(file)
    elif file.content_type in ["video/mp4", "video/quicktime"]:
        out = handlers.handle_video(file)
    else:
        raise HTTPException(status_code=400, detail="Incorrect data type")

    return out
