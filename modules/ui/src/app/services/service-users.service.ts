import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { environment } from '@environments/environment';
import { ServiceUserModel } from '@models/service-user.model';
import { MessagesService } from '@services/messages.service';
import { BehaviorSubject, Observable } from 'rxjs';
import { catchError, mergeMap, publishLast, refCount, tap } from 'rxjs/operators';

@Injectable({
    providedIn: 'root'
})
export class ServiceUsersService {

    private readonly users$: Observable<ServiceUserModel[] | undefined>;
    private usersSource = new BehaviorSubject<ServiceUserModel[] | undefined>(undefined);
    private usersCache$: Observable<ServiceUserModel[] | undefined> | undefined;
    private currentServiceId = '';

    constructor(
        private http: HttpClient,
        private messages: MessagesService,
    ) {
        this.users$ = this.usersSource.asObservable();
    }

    getUsers(serviceId: string): Observable<ServiceUserModel[] | undefined> {
        return this.loadUsers(serviceId).pipe(
            mergeMap(() => {
                return this.users$;
            }),
        );
    }

    createUser(serviceId: string, user: ServiceUserModel): Observable<ServiceUserModel> {
        return this.http
            .post<ServiceUserModel>(`${environment.apiUrl}/services/${serviceId}/users`, user)
            .pipe(
                tap(() => {
                    this.reloadUsers();
                }),
            );
    }

    editUser(serviceId: string, user: ServiceUserModel): Observable<ServiceUserModel> {
        return this.http
            .put<ServiceUserModel>(`${environment.apiUrl}/services/${serviceId}/users/${user.name}`, user)
            .pipe(
                tap(() => {
                    this.reloadUsers();
                }),
            );
    }

    deleteUser(serviceId: string, name: string): Observable<ServiceUserModel> {
        return this.http
            .delete<ServiceUserModel>(`${environment.apiUrl}/services/${serviceId}/users/${name}`)
            .pipe(
                tap(() => {
                    this.reloadUsers();
                }),
            );
    }

    changePassword(serviceId: string, userName: string | undefined, password: string) : Observable<ServiceUserModel> {
        return this.http
            .put<ServiceUserModel>(`${environment.apiUrl}/services/${serviceId}/users/${userName}`,
            {password, name: userName})
            .pipe();
    }
    private reloadUsers(): void {
        this.usersCache$ = undefined;
        if (!!this.currentServiceId) {
            this.loadUsers(this.currentServiceId).subscribe();
        }
    }

    private loadUsers(serviceId: string): Observable<ServiceUserModel[] | undefined> {
        if (serviceId !== this.currentServiceId) {
            this.usersCache$ = undefined;
            this.currentServiceId = serviceId;
        }

        if (!this.usersCache$) {
            this.usersCache$ = this.http
                .get<ServiceUserModel[] | undefined>(`${environment.apiUrl}/services/${serviceId}/users`)
                .pipe(
                    publishLast(),
                    refCount(),
                    catchError((err) => {
                        this.messages.error(err);
                        this.usersCache$ = undefined;
                        throw err;
                    }),
                );
        }
        return this.usersCache$.pipe(
            tap((users) => {
                this.usersSource.next(users);
            })
        );
    }
}
