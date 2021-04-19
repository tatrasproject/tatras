from .argocd_python_client import ArgoCD
from .config import ARGOCD_HOST, ARGOCD_USER, ARGOCD_PASSWORD
from typing import Optional
from fastapi import FastAPI


if __name__ == "__main__":

    client = ArgoCD(ARGOCD_HOST, ARGOCD_USER, ARGOCD_PASSWORD)
    app = FastAPI()


    @app.get("/")
    def read_root():
        return {"Hello": "World"}


    @app.get("/items/{item_id}")
    def read_item(item_id: int, q: Optional[str] = None):
        return {"item_id": item_id, "q": q}