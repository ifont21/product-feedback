import { BoardDetailDTO } from './board-types';
import { Inject, Injectable, Optional } from '@angular/core';
import { BasePath } from './base-path';
import { Observable } from 'rxjs';
import { HttpClient } from '@angular/common/http';

@Injectable({
  providedIn: 'root',
})
export class PFeedbackBoardSpecService {
  constructor(
    @Optional() @Inject(BasePath) private basePath: string,
    private http: HttpClient
  ) {}

  getBoardDetails(boardId: string): Observable<BoardDetailDTO> {
    return this.http.get<BoardDetailDTO>(`${this.basePath}/boards/${boardId}`);
  }
}
