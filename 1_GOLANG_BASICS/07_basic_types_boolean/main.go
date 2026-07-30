package main

import(
	"fmt"
)

func main(){
	isLogged := true;
	isAdmin := false;
	hasSubscription := false;

	// AND &&
	canOpenDashboard := isLogged && hasSubscription;

	// OR ||
	canDeletePost := isLogged || isAdmin;

	fmt.Println("Pode abrir o Dashboard: ", canOpenDashboard);

	fmt.Println("Pode deletar o post: ", canDeletePost);

}