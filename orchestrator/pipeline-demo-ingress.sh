# Define colors
BLUE='\e[34m'
NC='\033[0m' # No Color (Reset)
DEPLOYMENT_NAME="weber"

echo -e  "${BLUE}kubectl delete deployment "$DEPLOYMENT_NAME" ${NC}"
kubectl delete deployment "$DEPLOYMENT_NAME" 
echo -e  "${BLUE}kubectl delete service "$DEPLOYMENT_NAME" ${NC}"
kubectl delete service "$DEPLOYMENT_NAME" 
echo -e  "${BLUE}kubectl delete ingress "$DEPLOYMENT_NAME" ${NC}"
kubectl delete ingress "$DEPLOYMENT_NAME" 
echo -e  "${BLUE}kubectl apply -f demo.yaml ${NC}"
kubectl apply -f demo.yaml  
sleep 5
echo -e  "${BLUE}kubectl get pods,svc,ing ${NC}"
kubectl get pods,svc,ing 
echo -e  "${BLUE}kubectl describe pods -l app="$DEPLOYMENT_NAME" ${NC} "
kubectl describe pods -l app="$DEPLOYMENT_NAME" 
