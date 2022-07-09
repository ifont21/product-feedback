import { LandingPageComponent } from './pages/landing-page.component';
import { RouterModule, Routes } from '@angular/router';

const routes: Routes = [
  {
    path: '',
    pathMatch: 'full',
    redirectTo: 'home',
  },
  {
    path: 'home',
    component: LandingPageComponent,
  },
  {
    path: 'feedbacks',
    loadChildren: () =>
      import('./pages/board-feedbacks/board-feedback-bundle.module').then(
        (m) => m.BoardFeedbackBundleModule
      ),
  },
  {
    path: 'roadmap',
    loadChildren: () =>
      import('./pages/roadmap-view/roadmap-view-bundle.module').then(
        (m) => m.RoadMapViewBundleModule
      ),
  },
  {
    path: 'project-admin',
    loadChildren: () =>
      import('./pages/project-admin/project-admin-bundle.module').then(
        (m) => m.ProjectAdminBundleModule
      ),
  },
];

export const AppRoutingModule = RouterModule.forRoot(routes);
