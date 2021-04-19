## CONFIG FILE FOR APPLICATION. 
## ALL SERVER SETTINGS GO HERE.
import os
from dotenv import load_dotenv
load_dotenv()

# ArgoCD CLI tool doesn't always like 'https://' in front 
## of the ARGOCD_SERVER env var. For this reason we leave 
## it out of our app as well and append here.
ARGOCD_SERVER = os.environ.get("ARGOCD_SERVER", None)
ARGOCD_SERVER = f"https://{ARGOCD_SERVER}"

# The default username is 'admin'
# The default password is either the ArgoCD pod name,
## or the value found in 'argocd-initial-admin-secret'
## depending on your version of Argo.
ARGOCD_USER = os.environ.get("ARGOCD_USER", None)
ARGOCD_PASSWORD = os.environ.get("ARGOCD_PASSWORD", None)

# This is the list of *enabled* applications available 
## for end users to deploy. Simply comment a row out to 
## disable.
# The specifics of each are defined in their respective 
## yaml under 'app_config'.
SERVICE_CATALOG = {
    'APPS': [
        'postgresql',
        'pgadmin4',
        'nginx',
        'jupyter',
        'vscode',
    ]
}