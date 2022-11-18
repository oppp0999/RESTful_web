package com.example.firstboot

import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.PostMapping
import org.springframework.web.bind.annotation.RequestBody
import org.springframework.web.bind.annotation.RequestMapping
import org.springframework.web.bind.annotation.RestController
import javax.validation.Valid
import javax.validation.constraints.Size

data class ViewAccount(
    val id:Int,
    val name: String
)

data class CreateAccount(
    @field: Size(min=2, max=5)
    val name:String
)

@RestController
@RequestMapping("/accounts")
class AccountsController {

    @GetMapping("/")
    fun getAll(): Iterable<ViewAccount> =
        listOf(ViewAccount(id=1, name = "First"))

    @PostMapping("/")
    fun create(@Valid @RequestBody request: CreateAccount): ViewAccount =
        ViewAccount(id=2, name= request.name)

}