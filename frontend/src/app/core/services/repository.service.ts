import { Injectable } from '@angular/core';
import { AbstractControl, ValidationErrors } from '@angular/forms';

@Injectable({
  providedIn: 'root'
})
export class RepositoryService {
  async validateRepository(path: string): Promise<boolean> {
    await new Promise(resolve => setTimeout(resolve, 600));
    return true;
  }
  validateRepoPath(control: AbstractControl): ValidationErrors | null {
    // Simulate repository path validation
    const path = control.value;
    if (!path || path.length < 5) {
      return { invalidPath: true };
    }
    return null;

    // TODO: Implement real repository path validation
    // This could involve checking if the directory exists and contains a .git folder
  }
}

