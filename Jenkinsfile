pipeline {
  agent {
    docker {
      image 'golang:1.13'
    }
  }

  stages {
    stage('Test') {
      steps {
        sh """
        rm -f $WORKSPACE/test-results.{log,xml}
        mkdir -p influxql
        cp -a $WORKSPACE influxql

        cd influxql
        go get -v -t
        go test -v | tee $WORKSPACE/test-results.log
        """
      }

      post {
        always {
          sh """
          if [ -e test-results.log ]; then
            go get github.com/jstemmer/go-junit-report
            go-junit-report < $WORKSPACE/test-results.log > test-results.xml
          fi
          """
          junit "test-results.xml"
        }
      }
    }
  }
}
