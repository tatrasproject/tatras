import requests
import json

class APICall:
    ''' Generic api methods used across multiple classes
    '''
    def __init__(self, auth):
        self.auth = auth

    def get(self, item_name):
        ''' Retrieve a specific item from ArgoCD
        '''
        r = requests.get(f"{self.auth.url}/api/v1/{self.route}/{item_name}", headers=self.auth.headers)
        r_dict = json.loads(r.text)
        return r_dict    
    
    def all(self):
        ''' Return all items from ArgoCD
        '''
        r = requests.get(f"{self.auth.url}/api/v1/{self.route}", headers=self.auth.headers)
        r_dict = json.loads(r.text)
        return r_dict

    def post(self, data):
        ''' Create an item in ArgoCD.
            ARGS:
                data (dictionary) - find templates for 'data' in "argocd_tempaltes/"
            RETURNS:
                r_dict (dict) - response from ArgoCD
        '''
        r = requests.post(
            f"{self.auth.url}/api/v1/{self.route}", 
            data=json.dumps(data), headers=self.auth.headers
            )
        r_dict = json.loads(r.text)
        return r_dict

class Application(APICall):
    ''' Class for ArgoCD Application CRD.
        Used to make python client look nice 
        as well as for APICall method overrides. 
    '''
    route = 'applications'
        
class Project(APICall):
    ''' Class for ArgoCD Project CRD.
        Used to make python client look nice 
        as well as for APICall method overrides. 
    '''
    route = 'projects'
    
class ArgoCD:
    ''' Primary class used to establish a client to an ArgoCD server
    '''
    def __init__(self, url, username, password):
        self.url = url
        self.username = username
        self.password = password
        self.headers  = self.set_header()
        self.applications = Application(self)
        self.projects = Project(self)
        
    def get_token(self):
        data={
            'username': self.username,
            'password': self.password,
        }
        r = requests.post(
            f"{self.url}/api/v1/session",
            data=json.dumps(data),
        )
        r_dict = json.loads(r.text)
        return r_dict['token']
    
    def set_header(self):
        return {"Authorization": f"Bearer {self.get_token()}"}
        