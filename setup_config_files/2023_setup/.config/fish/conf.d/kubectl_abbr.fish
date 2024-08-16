# Basic kubectl Commands
abbr -a k "kubectl"
abbr -a kgp "kubectl get pods"
abbr -a kgn "kubectl get nodes"
abbr -a kgs "kubectl get services"
abbr -a kga "kubectl get all"
abbr -a kdp "kubectl describe pod"
abbr -a kdsvc "kubectl describe service"
abbr -a krp "kubectl rollout restart"
abbr -a kc "kubectl create"
abbr -a ka "kubectl apply -f"
abbr -a kd "kubectl delete -f"
abbr -a kex "kubectl exec -it"
abbr -a kl "kubectl logs"

# Namespace Management
abbr -a kns "kubectl config set-context --current --namespace="
abbr -a kgns "kubectl get namespaces"
abbr -a kn "kubectl get namespace"
abbr -a knc "kubectl create namespace"
abbr -a kdn "kubectl delete namespace"

# Context Management
abbr -a kctx "kubectl config use-context"
abbr -a kctxg "kubectl config get-contexts"
abbr -a kcctx "kubectl config current-context"

# Pods and Logs
abbr -a kgpo "kubectl get pods --output=wide"
abbr -a klf "kubectl logs -f"
abbr -a ke "kubectl exec -it"
abbr -a klog "kubectl logs"

# Deployment Management
abbr -a kgd "kubectl get deployments"
abbr -a kdpl "kubectl describe deployment"
abbr -a kdd "kubectl delete deployment"
abbr -a krsd "kubectl rollout status deployment"
abbr -a krd "kubectl rollout restart deployment"

# Services and Ingress
abbr -a kgs "kubectl get services"
abbr -a kdsvc "kubectl describe service"
abbr -a kdsvc "kubectl delete service"
abbr -a kgi "kubectl get ingress"
abbr -a kdi "kubectl describe ingress"
abbr -a kdei "kubectl delete ingress"

# ConfigMaps and Secrets
abbr -a kgcm "kubectl get configmaps"
abbr -a kgsec "kubectl get secrets"
abbr -a kdc "kubectl describe configmap"
abbr -a kdsec "kubectl describe secret"
abbr -a kdcm "kubectl delete configmap"
abbr -a kdsec "kubectl delete secret"

# Custom Abbreviations for Files
abbr -a kaf "kubectl apply -f file.yaml"
abbr -a kdf "kubectl delete -f file.yaml"
abbr -a kcf "kubectl create -f file.yaml"
abbr -a kf "kubectl apply -f"

# Additional Shortcuts
abbr -a krh "kubectl rollout history"
abbr -a kdr "kubectl describe"
abbr -a kw "kubectl get pods --watch"
abbr -a kdelp "kubectl delete pod"
abbr -a kre "kubectl replace --force -f"

