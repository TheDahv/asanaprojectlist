(function (angular) {
  angular.module('asanaProjectsViewer', [])
    .controller('ProjectsController',
    ['$scope', '$http', function ($scope, $http) {
      $http.get('/projects').success(function (projects) {
        $scope.projects = projects;
      });

      $scope.greeting = "Hello Asana!";
    }]);
})(window.angular);
