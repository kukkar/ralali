pipeline {
	agent {
		docker {
			image 'sanksons/go1.9-ubuntu:initial'
			args '-u 0:0'
			reuseNode true
		}
	}
    environment {
        // Some environment variables to be used later
        ARTIFACTPATH = "output"
        OUTPUT = "bundle.tar.gz"
		PROJECT_NAME = "reflorest-testapp"
		PROJECT_PATH = "github.com/sanksons"
	}
    stages {
	
	    //Take checkout from source control 
	    stage ('checkout') {
			steps {
				// mkdir where we will take checkout of source code.
				dir("sourcecode") {
					git branch: 'master', url: 'https://github.com/BoutiqaatREPO/getsetgo-testapp.git'
				}
            }
		}
		
		//Build project
	    stage('build') {
		    steps { 
				// copy contents from checkout out source code to gopath
				sh 'mkdir -p $GOPATH/src/$PROJECT_PATH'
				sh 'cp -r $WORKSPACE/sourcecode $GOPATH/src/$PROJECT_PATH/$PROJECT_NAME'	
				
				sh 'ls -lahrt $GOPATH/src/$PROJECT_PATH/$PROJECT_NAME'
					
				//Move inside project location and run deploy
				sh 'cd $GOPATH/src/$PROJECT_PATH/$PROJECT_NAME && reflorest deploy'
					 
			}
		}
		
		//Bundle data
		stage('bundle') {
		   steps {
		        //Copy the required files from bin to output folder
				sh 'rm -rf $ARTIFACTPATH'
                sh 'mkdir -p $ARTIFACTPATH'
				sh 'cp -r /gospace/bin/conf $ARTIFACTPATH/ && cp /gospace/bin/reflorest-testapp $ARTIFACTPATH/'
					
			    //Bundle archive.
     		    sh 'cd output && tar czf $OUTPUT *'
		        archiveArtifacts "${env.ARTIFACTPATH}/*"	  
			}
		}
	}
}
