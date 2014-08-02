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
    }]);
})(window.angular);
