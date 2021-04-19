from argocd_python_client import ArgoCD
from config import (
    ARGOCD_SERVER, ARGOCD_USER, ARGOCD_PASSWORD,
    SERVICE_CATALOG
)
from typing import Optional
from fastapi import FastAPI
import json


argocd = ArgoCD(ARGOCD_SERVER, ARGOCD_USER, ARGOCD_PASSWORD)
app = FastAPI()


@app.get("/")
def read_root():
    return {"Hello": "World"}

@app.post("/service_catalog/{item_name}")
def create_application(item_name: str, q: Optional[str] = None):
    ''' Endpoint for client to request the creation of an application. 
        This function checks config.py to ensure it's an authorized application,
        then uses the "argocd_templates/create-app.json" file as a template for the 
        ArgoCD RESTful API.
        Parameters are updated as needed, to make the resource unique to the client,
        and then sent to the ArgoCD server for creation. 
        ARGS:
          item_name (str) - Name of application to be created
        RETURNS:
          item_name
          argo_response
    '''
    if item_name not in SERVICE_CATALOG['APPS']:
        return {'status': "that service ain't allowed fam"}

    f = open('argocd_templates/create-app.json',)
    app = json.load(f)
    app['spec']['destination']['namespace'] = "test-namespace"
    app['spec']['source']['repoURL'] = "https://github.com/Extended-ERP/extendederp.git"
    app['spec']['source']['path'] = "kubernetes/helm/extendederp"
    app['spec']['source']['helm']['parameters'][0]['value'] = item_name
    app['metadata']['name'] = item_name

    print(app)

    r = argocd.applications.post(app)

    return {"item_name": item_name, "argo_response": r}