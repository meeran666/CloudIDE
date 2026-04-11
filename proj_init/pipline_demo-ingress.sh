# Define colors
BLUE='\e[34m'
NC='\033[0m' # No Color (Reset)

echo -e  "${BLUE}kubectl delete deployment service${NC}"
kubectl delete deployment service
echo -e  "${BLUE}kubectl delete service service${NC}"
kubectl delete service service
echo -e  "${BLUE}kubectl delete ingress service${NC}"
kubectl delete ingress service
echo -e  "${BLUE}kubectl apply -f demo.yaml${NC}"
kubectl apply -f demo.yaml 
sleep 5
echo -e  "${BLUE}kubectl get pods,svc,ing${NC}"
kubectl get pods,svc,ing
echo -e  "${BLUE}kubectl describe pods -l app=service${NC}"
kubectl describe pods -l app=service