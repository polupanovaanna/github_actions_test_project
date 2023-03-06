package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
	"project/util"
)

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalf(msg, err)
	}
}

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	//TODO ngrock ??

	CheckErr(err, "did not connect: %s")

	defer conn.Close()

	c := util.NewCommitDataClient(conn)

	masterHeadCommit := os.Getenv("INPUT_MASTERHEADHASH")
	//triggerCommit := os.Getenv("INPUT_TRIGGERHASH")
	diff := os.Getenv("INPUT_DIFF")

	response, err := c.Translate(context.Background(), &util.CommitInfo{HeadHash: masterHeadCommit, CommitDiff: diff})
	CheckErr(err, "Error when processing git info: %s")
	log.Printf("Response from server: %s", response)

	/*refIter, _ := r.References()

	refIter.ForEach(func(ref *plumbing.Reference) error {

		if ref.Name().IsRemote() {
			branchCommit, _ := r.CommitObject(ref.Hash())
			patch, _ := masterHeadCommit.Patch(branchCommit)

			response, err := c.Translate(context.Background(), &util.CommitInfo{HeadHash: masterHeadRef.Hash().String(),
				CommitDiff: patch.String()})
			CheckErr(err, "Error when translating info to server: %s")
			fmt.Println(response)
			fmt.Println("branch: ", patch.String())
		}
		return nil
	}) //iterating branches

	//log.Printf("Response from server: %s", response)*/

}
