import { CommonModule } from '@angular/common';
import { HttpClient, HttpClientModule } from '@angular/common/http';
import { Component, NgModule, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { environment } from '../environments/environment';

@Component({
  selector: 'app-projects',
  template: `
    <div>Project List over here</div>
    {{ projects$ | async | json }}
  `,
})
export class ProjectComponent implements OnInit {
  projects$!: Observable<unknown>;

  constructor(private http: HttpClient) {}

  ngOnInit(): void {
    this.projects$ = this.getProjects();
  }

  private getProjects(): Observable<unknown> {
    return this.http.get(`${environment.dev.serverUrl}/api/projects`);
  }
}

@NgModule({
  imports: [CommonModule, HttpClientModule],
  declarations: [ProjectComponent],
  exports: [ProjectComponent],
})
export class ProjectModule {}
