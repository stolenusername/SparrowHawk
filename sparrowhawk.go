package main

import "fmt"
import "os"

func main() {

	//Check for IOCs based on https://redcanary.com/blog/clipping-silver-sparrows-wings/
	iocs := []string{"~/Library/._insu", "/tmp/agent.sh", "/tmp/version.json", "/tmp/version.plist", "~/Library/Application Support/agent_updater/agent.sh", "/tmp/agent", "~/Library/Launchagents/agent.plist", "~/Library/Launchagents/init_agent.plist"}

	for i := 0; i < len(iocs); i++ {
		if _, err := os.Stat(iocs[i]); err == nil {
			fmt.Printf("ALERT!!!!!: File exists: "+iocs[i]+"\n"); 
			fmt.Printf("Go read the article at https://redcanary.com/blog/clipping-silver-sparrows-wings/\n") 
			fmt.Printf("   \n")
		  } else {
			fmt.Printf("File does not exist: "+iocs[i]+"\n");  
		  }
		}

}
