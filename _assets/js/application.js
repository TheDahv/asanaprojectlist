(function (angular) {
  var statusToSymbol = function (project) {
    switch (project.Status.toLowerCase()) {
      case 'r': return ':(';
      case 'y': return ':|';
      case 'g': return ':)';
      case 'unknown': return '?';
    }
  };

  angular.module('asanaProjectsViewer', [])
    .controller('ProjectsController',
    ['$scope', '$http', function ($scope, $http) {
      $http.get('/projects').success(function (projects) {
        $scope.projects = projects.map(function (p) {
          p.symbol = statusToSymbol(p);
          return p;
        });
      });

      $scope.greeting = "Hello Asana!";
    }]);
})(window.angular);
