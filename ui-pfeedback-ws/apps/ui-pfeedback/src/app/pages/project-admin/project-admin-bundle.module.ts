import { RouterModule } from '@angular/router';
import { Component, NgModule } from '@angular/core';

@Component({
  selector: 'app-project-admin-bundle',
  template: ` <div>Project Admin Page!</div> `,
})
export class ProjectAdminBundleComponent {}

const routes = [
  {
    path: '',
    component: ProjectAdminBundleComponent,
  },
];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  declarations: [ProjectAdminBundleComponent],
})
export class ProjectAdminBundleModule {}
