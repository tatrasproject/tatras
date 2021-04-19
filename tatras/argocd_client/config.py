import os
from dotenv import load_dotenv
load_dotenv()

ARGOCD_HOST = os.environ.get("ARGOCD_HOST", None)
ARGOCD_USER = os.environ.get("ARGOCD_USER", None)
ARGOCD_PASSWORD = os.environ.get("ARGOCD_PASSWORD", None)