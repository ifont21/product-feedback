import { LandingPageComponent } from './pages/landing-page.component';
import { RouterModule, Routes } from '@angular/router';
import { AuthGuard } from '@auth0/auth0-angular';

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
    canActivate: [AuthGuard],
  },
  {
    path: 'roadmap',
    loadChildren: () =>
      import('./pages/roadmap-view/roadmap-view-bundle.module').then(
        (m) => m.RoadMapViewBundleModule
      ),
    canActivate: [AuthGuard],
  },
  {
    path: 'project-admin',
    loadChildren: () =>
      import('./pages/project-admin/project-admin-bundle.module').then(
        (m) => m.ProjectAdminBundleModule
      ),
    canActivate: [AuthGuard],
  },
  {
    path: '**',
    redirectTo: 'home',
  },
];

export const AppRoutingModule = RouterModule.forRoot(routes);
