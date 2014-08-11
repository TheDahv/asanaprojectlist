(function (angular, _) {
  var statusToSymbol = function (project) {
    switch (project.Status.toLowerCase()) {
      case 'r': return ':(';
      case 'y': return ':|';
      case 'g': return ':)';
      case 'done': return ':D';
      case 'unknown': return '?';
    }
  };

  angular.module('asanaProjectsViewer', ['angularSlideables'])
    .controller('ProjectsController',
    ['$scope', '$http', function ($scope, $http) {
      $http.get('/projects').success(function (projects) {
        var positionCounter = 1;
        var prepareProject = function (p) {
          p.Symbol = statusToSymbol(p);
          p.Position = positionCounter;
          positionCounter++;
          return p;
        };

        $scope.projects = projects
          .filter(function (p) { return p.Status !== "Done"; })
          .map(prepareProject)

        $scope.completedProjects = projects
          .filter(function (p) { return p.Status === "Done"; })
          .map(prepareProject);
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

        if (!$scope.project.tasks) {
          $http.get('/projects/' + project.ID + '/tasks').success(function (tasks) {
            $scope.project.tasks = tasks;

            $scope.project.taskContributors = _.chain(tasks)
              .filter(function (task) {
                return task.Assignee && task.Assignee.Name.length > 0;
              })
              .groupBy(function (task) {
                return task.Assignee.Name;
              })
              .map(function (taskGroup) {
                return {
                  Name: taskGroup[0].Assignee.Name,
                  TaskCount: taskGroup.length
                };
              })
              .sortBy(function (taskGroup) {
                // Reverse order sort by task length
                return -1 * taskGroup.TaskCount;
              })
              .value();
          });
        }

        // Who knows how this actually gets into scope...
        event.preventDefault();
        event.stopPropagation();
      };
    }])
    .filter('outputIfMatch', function () {
      return function (input, rxPattern, output) {
        var rx = new RegExp(rxPattern, "i");
        input = input || '';

        if (rx.test(input)) {
          return output;
        } else {
          return '';
        }
      };
    });

})(window.angular, window._);
