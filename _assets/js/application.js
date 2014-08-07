(function (angular) {
  var statusToSymbol = function (project) {
    switch (project.Status.toLowerCase()) {
      case 'r': return ':(';
      case 'y': return ':|';
      case 'g': return ':)';
      case 'unknown': return '?';
    }
  };

  angular.module('asanaProjectsViewer', ['angularSlideables'])
    .controller('ProjectsController',
    ['$scope', '$http', function ($scope, $http) {
      $http.get('/projects').success(function (projects) {
        var positionCounter = 1;

        $scope.projects = projects.map(function (p) {
          p.Symbol = statusToSymbol(p);
          p.Position = positionCounter;
          positionCounter++;
          return p;
        });
      });

      $scope.orderProp = "Position";

      $scope.drawerActive = false;
      $scope.toggleDrawer = function () {
        $scope.drawerActive = !$scope.drawerActive;
      };
    }])
    .controller('ProjectItemController',
    ['$scope', '$http', function ($scope, $http) {
      $scope.toggleDetails = function (project) {
        $scope.project.showDetails = !$scope.project.showDetails;

        if (!$scope.project.details) {
          $http.get('/projects/' + project.ID).success(function (details) {
            $scope.project.details = details;
          });
        }

        if (false && !$scope.project.tasks) {
          $http.get('/projects/' + project.ID + '/tasks').success(function (tasks) {
            $scope.project.tasks = tasks;
          });
        }

        return false;
      };
    }])
    .filter('ifMatch', function () {
      return function (input, rxPattern) {
        var rx = new RegExp(rxPattern, "i");
        input = input || '';

        if (rx.test(input)) {
          return input;
        } else {
          return '';
        }
      };
    });

})(window.angular);
