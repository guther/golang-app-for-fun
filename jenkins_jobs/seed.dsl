pipelineJob("CI") {
    logRotator{
        numToKeep 30
    }
    definition {
        cps {
            sandbox()
            script("""
                node {
                    stage("Checkout"){
                        echo 'Starting Checkout'
                        try{
                            sh 'git clone -b dev git@github.com:guther/hostgator-challenge.git'
                        }
                        catch(ex){
                            sh 'cd hostgator-challenge && pwd && git checkout'
                        }
                    }
                    stage("Build + Unit tests"){
                        sh 'cd hostgator-challenge'
                        echo 'Starting Unit Tests'
                        sh 'docker exec web_container go test /go/src/tests/ > unit_tests.log'
                    }
                    
                    stage("Archiving Reports"){
                        echo "Goodbye world"
                       
                    }
                }
            """.stripIndent())
        }
    }
}



