Keeva Operator (speak2jc/k-op)

The purpose of this project is to demonstrate the use of the k8s operator-sdk in conjunction with a custom resource definition (CRD).

How it works
The operator-sdk creates all of the scaffolding for an operator and allows "api" and "controller" scaffold code to be generated.
Think of the api as a set of structs representing the CRD you are about the create and the controller then watches the target cluster for instances of
this CR and reconciles the k8s cluster according to the instructions in the controller.

Target Cluster
The taregt cluster can be any k8s cluster you can access - for testing I use minikube locally. 
Before an instance of given CR can be created, the CRD must be registered with the cluster. The operator-sdk also creates a CRD (in deploy folder)
which yuo can use to get started - you will obviously need to add any new fields as you design your actual CR.
Registering can be done by simplying applying the CRD to the cluster (or can also be done programmatically via client-go more about that later).

Note that the controller can run in-cluster (in a pod) or outside (locally on your Macbook - as long as it can connect via config found 
inside .kube/config) - typically a snippet of code is used to check fist for in-cluster and then fall back on outside-cluster.

Who creates the CR?
The secondary purpose of this project is to test out a pattern whereby CR/CRD and controllers are defined in this project, but the generated objects 
used in a different project. The second project demonstrates how to use client-go for CRUD operations with a CRD.
The client-go library provides a way in golang to interact directly with the k8s cluster (as opposed to using shell to call kubectl). Client 
objects are strongly typed which means that the client needs to knwo the structure of the objects it handles - the default client from k8s is aware of 
all of the standard objects (deployments, pods etc.) but has no knowledge of any new objects defined in CRDs.

Note: a valid alternative is still to use "kubectl apply" in a shell spawned from the appropriate point in the second project - this has pros and cons. 
It is not very integrated insofar as you would maintain a template YAML file and perform substitution into a tmp dir before applying using kubectl.
On the other hand, all the second project needs is a valid copy of the CR structure.

Creating a custom client
There are three ways to create a client. You could us a dynamic client - there is an example in speak2jc/cr-manager. The problem here is that the objects 
are maps and therefore untyped. As a result you need mappers to convert between your domain objects (the structs representing your CR) and generic maps; 
furthermore the right side of the map is an interface, therefore so generic that you need to use explicit casting, which is ugly and risky.
The second method would be to hand-code a typed client - let's forget that idea.

Thirdly, you can generate client code in this project, just after you've run operator-sdk - you then access this code in the second project and use in 
exactly the same way as with the k8s default client. This is a great approach, which makes it surprising that it is very poorly documented anywhere.

Steps
- Assuming we use go modules, add a dependency on code-generation into go.mod as shown in this project (use same version for now)
- Run go mod vendor to grab the code generator into the vendor folder (presumably an alternative would be to copy the generate-groups.go file and store it within the project and then skip the vendor step in future)
- Annotate the keevakind_types file with +genclient   - the tricky part here is to know which needs annotating - you should ONLY annotate the top-level object, else lots of errors.
- Run the ./hack/update-codegen.sh file - this is just a convenient way of re-running generate-groups with parameters.
- You should see a pkg/generated folder appear and inside will be the typed client files.
- Unless you need the vendor folder for something else (note: certain versions of operator-sdk do), delete it and save headaches later when the code is imported into second project.
- Optionally rebuild code using operator-build to check that all still works 

Observations
- If you annotate too many files in types files, sometimes you get ObjectMeta errors during generation and other times it all works but you get issues in the Lists (discovered often only when importing to second project)
- The hack code is just a hack - double check that it matches your project structure (e.g. number of cd ../../)
- If you don't delete vendor folder, then when importing into second project when developing locally, the client will still try to refer to files in here - as vendored files cannot be imported (and shouldn't) you get reams of error messages.
- In contrast to corporate setup, changes in this project seem to be immediately effective in the second project even if not yet pushed to github.com (this is really good!) 


 
