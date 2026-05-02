kubectl exec -it ba02bd3f-1a9d-4da7-a58f-bbf93a05e83f-86c6c65bbc-q2w2v -- /bin/sh
kubectl get pods,svc,ing -A
kubectl port-forward meer-54d83146-cdf5-41a2-9da2-3e167ee9db32-866db79c44-64tkh -n ide  3011:3006
kubectl port-forward service/meeran-62af68f9-1552-49cd-b3b8-80504bae2d4f 3011:3006
docker run -it file_accpeter /bin/sh
docker build -t file_accpeter .
kubectl delete ingress --all -n my-namespace
kubectl get all -n default
kubectl logs -n ide meer-7a0f9368-259e-437e-937f-58d1159a6720-78bc5d4bb4-58nfb